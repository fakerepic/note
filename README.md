# note

## TODO(backend)

### app gateway

- [ ] Add auto migration for schema, email sender, admin users, etc.
- [ ] Configure S3 bucket for user uploaded attachments.

### ai service

- [ ] Configure the cache path of embedding model.
- [ ] Implement `GET /ai/last_refresh_time`.
- [ ] Schedule auto refresh for user.
- [ ] Use thread pool to handle requests.
- [ ] Move to rust version of pgvector: `pgvecto.rs`
