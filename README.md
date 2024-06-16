# note

## Deploy with docker

Before you run the following command, you need to put `.env.prod` file in the root directory of the project. We have provided the `.env.example` file for you to refer to.

```bash
docker compose up -d
```

### How to configure the `.env.prod` file based on the `.env.example` file

You need to change the following variables:

- `APP_PUBLIC_URL`: The public URL of the application. It is essential for the email verification link. (It is unnecessary to change the default value `http://127.0.0.1:8090` if you just want to test the application locally.)

- `ADMIN_EMAIL`, `ADMIN_PASSWORD`: The email and password of the administrator account (of pocketbase). The email should be a valid email address and the password should be at least 10 characters long and contain at least 1 upper case letter.

- `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASSWORD`: The SMTP server configuration for sending emails. For example, you can use Tsinghua's SMTP server:

  > ```bash
  > SMTP_HOST=mails.tsinghua.edu.cn
  > SMTP_PORT=465
  > SMTP_USER=XXXXXX@mails.tsinghua.edu.cn # Your Tsinghua email address
  > SMTP_PASSWORD=YOUR_TSINGHUA_EMAIL_TOKEN # You can generate an token in your tsinghua email account settings.
  > ```

- `COUCHDB_USER`, `COUCHDB_PASSWORD`: The username and password of the CouchDB database. Don't leave them empty.

- `TOGETHER_API_KEY`: The API key of Together AI. You can get it for free from the [Together AI website](https://www.together.ai/)
