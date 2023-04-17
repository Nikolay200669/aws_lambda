# aws_lambda

### create build & zip file for upload to AWS lambda

```shell
make all
```


### test
```shell
curl --location 'https://url_to_you_aws_lambda_function.on.aws/?name=my_test' \
--header 'Content-Type: text/plain' \
--header 'Authorization: Bearer >>>>> test_token <<<<<' \
--data '{
    "integer":123
    "test_string":"Lorem ipsum for test"
}'
```
