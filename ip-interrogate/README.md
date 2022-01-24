# ip-interrogate
# https://0ifursd48i.execute-api.us-east-2.amazonaws.com/dev/hello <-- current live API
A tool to interrogate a single IP deployed via terraform on AWS lambda written in GO
Sources cited:
https://www.youtube.com/watch?v=Wd-xhOKkMcU
https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html#api-gateway-overview-aws-backbone
https://learn.hashicorp.com/tutorials/terraform/lambda-api-gateway
https://serverless-stack.com/examples/how-to-create-a-rest-api-in-golang-with-serverless.html
https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started
https://stackoverflow.com/questions/57054136/how-to-resolve-this-error-archiving-directory-could-not-archive-missing-direct
https://github.com/jagonzalr/go-lambda-terraform-setup/blob/main/infrastructure/main.tf
https://www.qloudx.com/understanding-the-terraform-resources-that-create-an-aws-api-gateway-rest-api/
https://www.rfc-editor.org/rfc/rfc7159.html
https://tpaschalis.github.io/golang-aws-lambda-getting-started/
# https://0ifursd48i.execute-api.us-east-2.amazonaws.com/dev/hello <-- current live API
# the API currently only suppoerts POST and OPTIONS requests, I had other versions, that did other things 'correctly' (printing an arbitrary string on a successful request)
# This is the closest I got for the intended function.
# a quick note on why I left the API in this state
Loading response...
Invocation result for arn:aws:lambda:us-east-2:786755870367:function:ip-interrogate
Logs:
START RequestId: 60d746aa-b626-45d7-9ef5-348b3310cb1a Version: $LATEST
# **2022/01/24 02:36:11 whois: domain is empty** <-- I got really close to a working API call, but as the code is today, I couldn't populate my 'first' variable remotely
2022/01/24 02:36:11 unexpected EOF
2022/01/24 02:36:11 unexpected EOF
END RequestId: 60d746aa-b626-45d7-9ef5-348b3310cb1a
REPORT RequestId: 60d746aa-b626-45d7-9ef5-348b3310cb1a	Duration: 1.92 ms	Billed Duration: 2 ms	Memory Size: 128 MB	Max Memory Used: 29 MB	Init Duration: 77.09 ms	
RequestId: 60d746aa-b626-45d7-9ef5-348b3310cb1a Error: Runtime exited with error: exit status 1
Runtime.ExitError


Payload:
{"errorMessage":"RequestId: 60d746aa-b626-45d7-9ef5-348b3310cb1a Error: Runtime exited with error: exit status 1","errorType":"Runtime.ExitError"}


This was a pretty cool way to spend a weekend. Looking forward to talking with you about it and meeting the team.
-Matt