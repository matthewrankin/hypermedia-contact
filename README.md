# Hypermedia Contacts

Example application developed as part of reading _Hypermedia Systems_.

## Development

To spin up the Docker compose file, migrate the database, and then seed the
database use:

```bash
$ make up
$ make dbup
$ make dbseed
```

### Development Dependencies

1. Docker
2. go migrate
