#!/usr/bin/env bash

while getopts ":c:n:s:m" opt; do
  case $opt in
    c ) conf_dir=$OPTARG;;
    n ) db_name=$OPTARG;;
    s ) server_port=$OPTARG;;
    m ) mqtt_conf="true";;
    : ) echo "Missing argument for option -$OPTARG"; exit 1;;
    \?) echo "Unknown option -$OPTARG"; exit 1;;
  esac
done

[ -z "$conf_dir" ] && mkdir -p "$HOME/.config/moody" && conf_dir="$HOME/.confg/moody"
[ -z "$db_name" ] && db_name="activity"
[ -z "$server_port" ] && server_port=7000

mqtt_conf='{
    "mqtt": {
        "host": "'$HOSTNAME'",
        "port": 8883,
        "tlsEnabled": true,
        "dataTopic": [
            "moody/service/+",
            "moody/actserver"
        ],
        "pubTopics": [
            "moody/actuator/mode",
            "moody/actuator/situation"
        ]
    }
}'

conf='{
    "conf_dir": "'$conf_dir'",
    "db_name": "'$db_name'"
    "server_port": "'$server_port'",
    '$mqtt_conf'
}'

mkdir -p data
echo "$conf" > data/conf.json