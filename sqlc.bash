
# ENV_PATH="./build"
# ENV_QUERIES="./query"
# ENV_SCHEMA="./schema"

BASE="sqlc-temp"

CONTAINER="sqlcc"

sqlc1(){
  rm -rf sqlc-temp
  mkdir $BASE
  mkdir $BASE/query
  mkdir $BASE/schema
  
  echo '
version: "1"
packages:
  - name: "db"
    path: "./build/"
    queries: "./query/"
    schema: "./schema/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: false
    emit_interface: false
    emit_exact_table_names: false  
' >> $BASE/sqlc.yaml

  echo "把檔案放入$BASE資料夾"
}

execc=exec
sqlc2(){
  docker $execc $CONTAINER sh -c 'rm -rf ./*'
  docker cp $BASE/. $CONTAINER:/app/sqlc-temp
  docker $execc $CONTAINER sqlc generate
  mkdir $BASE/build
  docker cp $CONTAINER:/app/sqlc-temp ./
}