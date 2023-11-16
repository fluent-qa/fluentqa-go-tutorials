docker run -d -p 3000:3000 \
    -e PREST_PG_URL=postgres://postgres:changeit@localhossst:7432/test_hub \
    -e PREST_DEBUG=true \
    prest/prest:v1