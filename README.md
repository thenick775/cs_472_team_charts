# cs_472_team_charts

![](https://github.com/thenick775/cs_472_team_charts/blob/main/demo.gif)

[Try it live](https://rawcdn.githack.com/thenick775/cs_472_team_charts/7497c37b2748ee99a2fdb5dd865770eb1d4bbee2/line.html)

This is a short program written in go to make generating team reporting charts more flexible, faster, and interactive

Generation instructions:

- use `go run main.go` to generate the `line.html` file

- open the `line.html` file in any browser

Update instructions:

- to update team commits, open main.go and adjust the values in the valsc variable instance of []TeamData

- to update team tasks, open main.go and adjust the values in the valst variable instance of []TeamData

- to adjust the names and order that names apear, adjust the order variable instance of []String

Features:

 -Download each chart as an image

- View raw data input in graphic

- Select which lines are displayed to inspect max value in the tooltip
