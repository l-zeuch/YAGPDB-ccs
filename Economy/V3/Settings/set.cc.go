{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(set|config(ure)?)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}

{{/* Configures economy settings */}}

{{/*
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$sdict := (dbGet 0 "EconomySettings").Value}}
{{$perms := split (index (split (exec "viewperms" ) "\n" ) 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
    {{with .CmdArgs}}
        {{with (dbGet 0 "EconomySettings")}}
            {{$a := sdict .Value}}
            {{if index $.CmdArgs 0}}
                {{$setting := (lower (index $.CmdArgs 0))}}
                {{if eq $setting "max"}}
                    {{$min := (toInt $a.min)}}
                    {{$symbol := $a.symbol}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$max := (index $.CmdArgs 1)}}
                        {{if (toInt $max)}}
                            {{if lt (toInt $max) (toInt $min)}}
                                {{$embed.Set "description" (print "You cannot set `" $setting "` to a value below `min`\n`min` is set to `" $min "`")}}
                                {{$embed.Set "color" $errorColor}}
                            {{else}}
                                {{$embed.Set "description" (print "You set `" $setting "` to " $symbol $max)}}
                                {{$embed.Set "color" $successColor}}
                                {{$sdict.Set "max" $max }}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else if eq $setting "min"}}
                    {{$max := (toInt $a.max)}}
                    {{$symbol := $a.symbol}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$min := (index $.CmdArgs 1)}}
                        {{if (toInt $min)}}
                            {{if gt (toInt $min) (toInt $max)}}
                                {{$embed.Set "description" (print "You cannot set `" $setting "` to a value above `max`\n`max` is set to `" $max "`")}}
                                {{$embed.Set "color" $errorColor}}
                            {{else}}
                                {{$embed.Set "description" (print "You set `" $setting "` to " $symbol $min)}}
                                {{$embed.Set "color" $successColor}}
                                {{$sdict.Set "min" $min }}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else if eq $setting "startbalance"}}
                    {{$symbol := $a.symbol}}
                    {{$oldStartBalance := $a.startBalance }}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$startBalance := (index $.CmdArgs 1)}}
                        {{if (toInt $startBalance)}}
                            {{$embed.Set "description" (print "You set `" $setting "` to " $symbol $startBalance " from " $oldStartBalance)}}
                            {{$embed.Set "color" $successColor}}
                            {{$sdict.Set "startBalance" $startBalance }}
                        {{else}}
                        {{$embed.Set "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
                        {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else if eq $setting "symbol"}}
                    {{if gt (len $.CmdArgs) 1}}
                        {{$symbol := (index $.CmdArgs 1)}}
                        {{$output := ""}}
                        {{if (reFind `(<a?:[A-z+]+\:\d{17,19}>)` $symbol)}}
                            {{$output = $symbol}}
                        {{else}}
                            {{$output = (print "`" $symbol "`")}}
                        {{end}}
                            {{$embed.Set "description" (print "You set the server currency symbol to " $output )}}
                            {{$embed.Set "color" $successColor}}
                        {{$sdict.Set "symbol" $symbol }}
                    {{else}}
                        {{$embed.Set "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:String>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else if eq $setting "default"}}
                    {{$embed.Set "description" (print "Set the `EconomySettings` to the default values")}}
                    {{$embed.Set "color" $successColor}}
                    {{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "symbol" "£" "startBalance" 200)}}
                    {{$staffBadge := "<:Badge_Staff:923744973118119946>"}}
                    {{$betaBadge := "<:Badge_Beta:923744973285892106>"}}
                    {{dbSet 0 "EconomyBadges" (sdict "staffBadge" $staffBadge "betaBadge" $betaBadge)}}
                {{else}}
                    {{$embed.Set "description" (print "No valid setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int>`\nAvailable settings: `max`, `min`, `startbalance`, `symbol`\nTo set it with the default settings `" $.Cmd " default`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "No database found.\nPlease set it up with the default values using `" $.Cmd " default`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int>`\nAvailable settings: `max`, `min`, `startbalance`, `symbol`\nTo set it with the default settings `" $.Cmd " default`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}
{{dbSet 0 "EconomySettings" $sdict}}