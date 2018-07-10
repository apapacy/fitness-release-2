-- if client IP is in whitelist, pass
local whitelist = ngx.shared.nla_whitelist
in_whitelist = whitelist:get(ngx.var.remote_addr)
if in_whitelist then
    return
end

-- HTTP headers
local headers = ngx.req.get_headers();

-- wp ddos
if type(headers["User-Agent"]) ~= "string"
    or headers["User-Agent"] == ""
    or ngx.re.find(headers["User-Agent"], "^WordPress", "ioj") then
    ngx.log(ngx.ERR, "ddos")
    ngx.exit(444)
    return
end

-- Это мобильные клинеты или 1с
if ngx.re.find(headers["User-Agent"], "^okhttp", "ioj")
    or ngx.re.find(headers["User-Agent"], "^android;")
    or ngx.re.find(headers["User-Agent"], "^Dalvik", "ioj")
    or ngx.re.find(headers["User-Agent"], "^Aura-iOS", "ioj")
    or ngx.re.find(headers["User-Agent"], "^ios;", "ioj")
    or ngx.re.find(headers["User-Agent"], "^1C\\+Enterprise", "ioj")  then
    return
end

local banlist = ngx.shared.nla_banlist
local search_bot = "search:bot:count:request:per:10:s"
if ngx.re.find(headers["User-Agent"], "Google Page Speed Insights|Googlebot|baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator", "ioj") then
   local count, err = banlist:incr(search_bot, 1)
    if not count then
        banlist:set(search_bot, 1, 10)
        count = 1
    end
    if count >= 50 then
        if count == 50 then
            ngx.log(ngx.ERR, "bot banned")
        end
        ngx.exit(444)
        return
    end
    return
end


-- cookies
local cookie = require("cookie")
local cookies = cookie.get()

-- global shared dict
local config = ngx.shared.nla_config
local req_count = ngx.shared.nla_req_count
local net_count = ngx.shared.nla_net_count

-- config options
local COOKIE_NAME = config:get("cookie_name")
local COOKIE_SID = config:get("cookie_sid")
local COOKIE_KEY = config:get("cookie_key")
local REQUESTS_PER_SECOND = config:get("requests_per_second")

-- identify if request is page or resource
if ngx.re.find(ngx.var.uri, "\\/.*?\\.[a-z]+($|\\?|#)", "ioj")
    and not ngx.re.find(ngx.var.uri, "\\/.*?\\.(html|htm|php|py|pl|asp|aspx|ashx)($|\\?|#)", "ioj") then
    ngx.ctx.nla_rtype = "resource"
else
    ngx.ctx.nla_rtype = "page"
end

-- init random
math.randomseed(os.time() + os.clock())

-- get or set client seed
local sid
if cookies[COOKIE_SID] == nil then
    sid = math.random(100000000000, 999999999999)
else
    sid = cookies[COOKIE_SID]
end

-- session tokens
local user_id = ngx.md5(ngx.var.remote_addr .. ngx.var.hostname .. (headers["User-Agent"] or "") .. COOKIE_KEY .. sid)
local network_id = ngx.md5(ngx.var.remote_addr .. ngx.var.hostname .. (headers["User-Agent"] or ""))

local count, err = req_count:incr(user_id, 1)
if not count then
    req_count:set(user_id, 1, 10)
    count = 1
end

-- counter from ip
if not cookies[COOKIE_NAME] then
    local count, err = net_count:incr(network_id, 1)
    if not count then
        net_count:set(network_id, 1, 3600 + math.random(0, 600))
        count = 1
    end
    if count >= 1024 then
        if count == 1024 then
            ngx.log(ngx.ERR, "client banned by network")
        end
        ngx.exit(444)
        return
    end
    cookie.challenge(COOKIE_NAME, user_id, COOKIE_SID, sid)
    return
end

-- counter from sid
if cookies[COOKIE_NAME] ~= user_id then
    local count, err = banlist:incr(cookies[COOKIE_NAME], 1)
    if not count then
        banlist:set(cookies[COOKIE_NAME], 1, 3600 + math.random(0 , 600))
        count = 1
    end
    if count >= 1024 then
        if count == 1024 then
            ngx.log(ngx.ERR, "client banned by bad sid")
        end
        ngx.exit(444)
        return
    end
    cookie.challenge(COOKIE_NAME, user_id, COOKIE_SID, sid)
    return
end

-- counter from sid
if count >= 512 then
    local count, err = banlist:incr(cookies[COOKIE_NAME], 1)
    if not count then
        banlist:set(cookies[COOKIE_NAME], 1, 3600 + math.random(0, 600))
        count = 1
    end
    if count >= 512 then
        if count == 512 then
            ngx.log(ngx.ERR, "client banned by retry")
        end
        ngx.exit(444)
        return
    end
    cookie.challenge(COOKIE_NAME, user_id, COOKIE_SID, sid)
    return
end

if (count > REQUESTS_PER_SECOND) then
    cookie.challenge(COOKIE_NAME, user_id, COOKIE_SID, sid)
    return
end
