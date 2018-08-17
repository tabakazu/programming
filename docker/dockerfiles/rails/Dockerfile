FROM node:8.11.3-alpine as nodejs
FROM ruby:2.3.7-alpine

RUN apk add --no-cache libgcc build-base libxml2-dev libxslt-dev zlib-dev postgresql-dev \
  && mkdir /opt && mkdir /usr/local/lib/node_modules
COPY --from=nodejs /usr/local/bin/node /usr/local/bin/
COPY --from=nodejs /usr/local/lib/node_modules/npm /usr/local/lib/node_modules/npm
COPY --from=nodejs /opt/yarn-v1.6.0 /opt/yarn-v1.6.0
RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm \
  && ln -s /opt/yarn-v1.6.0/bin/yarn /usr/local/bin/yarn \
  && ln -s /opt/yarn-v1.6.0/bin/yarnpkg /usr/local/bin/yarnpkg
RUN mkdir /webapp
WORKDIR /webapp
COPY Gemfile /webapp/Gemfile
COPY Gemfile.lock /webapp/Gemfile.lock
RUN bundle install
COPY . /webapp
