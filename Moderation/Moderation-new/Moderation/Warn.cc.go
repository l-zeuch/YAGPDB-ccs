{{/*
        Case system by Maverick Wolf (549820835230253060)
        Made & modified by Ranger (765316548516380732)

    **Tools & Util > Moderation > Warn > Warn DM**
©️ Dynamic 2021
MIT License
*/}}


{{/*        Notes
    This command allows for use with the `BannedWords` CC (Unsure who made it), if you haven't added it, and want to implement it DM me
    Please be aware that even though the original custom commands had the time, I removed it from this for NOW. They may be added back if I can get it to show the current time, not the time in GMT.
*/}}

{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}} {{/* Modlog channel */}}
{{$dm := 1}} {{/* change to 0 if you don't wanna DM the offender about the moderation action */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
	{{$ext := "webp"}}
	{{if eq (slice .Guild.Icon 0 2) "a_"}} {{$ext = "gif"}} {{end}}
	{{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{$bannedWords := ""}}

{{if (dbGet 0 "banned words")}}
    {{$bannedWords = reReplace `\A` (toString (dbGet 0 "banned words").Value) "("}}
    {{$bannedWords = reReplace `\z` $bannedWords ")"}}
    {{$bannedWords = reReplace `\s` $bannedWords "|"}}
{{end}}

{{$case_number := (toInt (dbIncr 77 "cv" 1))}}
{{$action := .ModAction.Prefix}}
{{$a := 0}}

{{if eq $action "Muted" "Unmuted"}}
    {{$a = (sub (len .ModAction.Prefix) 1)}}
    {{else}}
    {{$a = (sub (len .ModAction.Prefix) 2)}}
{{end}}

{{$title := (slice .ModAction.Prefix 0 $a)}}
{{$id := .User.ID}}
{{$channel := $LogChannel}}

{{$reason := ""}}
{{if .Reason}}
    {{if and (reFind "(?i)word blacklist" .Reason) (dbGet 0 "banned words")}}
        {{$reason = (print "Sending  word ||" (reFind $bannedWords (lower .Message.Content)) "|| is forbidden")}}
        {{else if or (not (reFind "(?i)word blacklist" .Reason)) (not (dbGet 0 "banned words"))}}
        {{if reFind `Automoderator:` .Reason}}
            {{$reason = (reReplace `Triggered rule:\s` (reReplace `Automoderator:\s` .Reason "") "")}}
            {{else}}
            {{$reason = .Reason}}
        {{end}}
    {{end}}
{{end}}

{{/* log & DM messages */}}
{{if $dm}}
    {{$WarnDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Warn`\n**Reason: **" $reason)
            "thumbnail" (sdict "url" $icon)
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
    {{sendDM $WarnDM}}
{{end}}

{{$x := sendMessageRetID $LogChannel (cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")")) "description" (print "<:TextChannel:800978104105304065> **Case number:** " $case_number "\n<:Management:788937280508657694> **Who:** " .User.Mention " `(ID " .User.ID ")`\n<:Metadata:788937280508657664> **Action:** `Warn`\n<:Assetlibrary:788937280554926091> **Channel:** <#" .Channel.ID ">\n<:Manifest:788937280579698728> **Reason:** " $reason "\n<:Edit:800978104272683038> **Message Logs:** [Click Here](" (execAdmin "logs") ")")
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "color" 16556627
            )}}

{{$Response := sendMessageRetID nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Case type: Warning"))
            "description" (print .Author.Mention " Has successfully warned " .User.Mention " :thumbsup:")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            )}}
{{deleteMessage nil $Response 5}}

{{/*for viewcase*/}}
{{dbSet $case_number "viewcase" (sdict "name" .Author.Username "warnname" .User.Username "avatar" (.Author.AvatarURL "512") "reason" $reason "userid" $id "action" (.ModAction.Prefix) "channel" $channel "msgid" $x "userdiscrim" .User.Discriminator)}}

{{/*for per user case viewing*/}}
{{dbSet $case_number $id (print "Case # **" $case_number "**\t\t**| " $title " Reason:** `" $reason "`")}}

{{/* for delete case*/}}
{{dbSet $case_number "userid" (str $id)}}