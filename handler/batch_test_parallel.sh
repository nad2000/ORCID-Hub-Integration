export URL=${1:-https://415mdw939a.execute-api.ap-southeast-2.amazonaws.com/prod/v1/enqueue}

xargs -P 100 -L 1 -i -t curl -v -H "Content-Type: application/json" -d "{\"subject\":{}}" $URL <<EOF
4306445
484378182
477579437
208013283
987654321
4306445
484378182
477579437
208013283
987654321
8524255
350622514
4306445
8524255
350622514
4306445
477579437
208013283
987654321
4306445
484378182
477579437
208013283
987654321
8524255
350622514
4306445
4306445
484378182
477579437
208013283
987654321
8524255
350622514
4306445
EOF
