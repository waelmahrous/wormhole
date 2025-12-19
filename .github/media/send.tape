Output send.gif

Require echo
Require wormhole

Set WindowBar Colorful
Set Shell "zsh"
Set FontSize 25

Set Margin 10
Set MarginFill "#FFAAFF"
Set BorderRadius 10

Set Width	1200
Set Height	600

# Showtime!
Type "wormhole --id demo --status" Sleep 300ms Enter
Sleep 800ms

Type "wormhole --id demo send a_file" Sleep 300ms Enter
Sleep 2s

Sleep 5s
