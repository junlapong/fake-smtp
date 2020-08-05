# fake-smtp

- [FakeSMTP](https://github.com/junlapong/FakeSMTP) java, desktop

```
java -jar jar/fakeSMTP-2.1-SNAPSHOT.jar -p 2525 -s
```

- [fake-smtp-server](https://github.com/gessnerfl/fake-smtp-server) java, web

```
java -jar -Dfakesmtp.port=2525 -Dserver.port=8025 jar/fake-smtp-server-1.5.0.jar
```

- [fake-smtp-server](https://github.com/ReachFive/fake-smtp-server) javascript, web

```
# npm install -g fake-smtp-server
fake-smtp-server -s 2525 -h 8025
```

- [MailHog](https://github.com/mailhog/MailHog) go

```
brew install mailhog

# or

go get github.com/mailhog/MailHog

MailHog -smtp-bind-addr 0.0.0.0:2525
```

## Web

[http://localhost:8025](http://localhost:8025/)

## SMTP

localhost:2525

## Test Send Mail

```
# cd test-send-mail
# go mod tidy

go run main.go

Try sending mail...
Mail sent without error
```

## Others

- [Hermes](https://github.com/matcornic/hermes) - package that generates clean, responsive HTML e-mails
- [Different ways to send an email with Golang](https://www.loginradius.com/engineering/blog/sending-emails-with-golang/)

