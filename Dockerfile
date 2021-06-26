# syntax=docker/dockerfile:1
FROM redis:latest
EXPOSE 6379
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]