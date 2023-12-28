CREATE TABLE user (
  id SERIAL PRIMARY KEY,
  username VARCHAR(15) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  complete_name VARCHAR(50) NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE tweet (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  content VARCHAR(160) NOT NULL,
  parent_id INTEGER,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_tweet_user FOREIGN KEY (user_id) REFERENCES user(id),
  CONSTRAINT fk_tweet_tweet FOREIGN KEY (parent_id) REFERENCES tweet(id)
);

CREATE TABLE user_following (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  following_user_id INTEGER NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT unique_following UNIQUE (user_id, following_user_id),
  CONSTRAINT fk_user_following_user FOREIGN KEY (user_id) REFERENCES user(id),
  CONSTRAINT fk_user_following_following_user FOREIGN KEY (following_user_id) REFERENCES user(id)
);

CREATE TABLE tweet_map_child_tweet (
  id SERIAL PRIMARY KEY,
  tweet_id INTEGER NOT NULL,
  child_tweet_id INTEGER NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_tweet_map_child_tweet_tweet FOREIGN KEY (tweet_id) REFERENCES tweet(id),
  CONSTRAINT fk_tweet_map_child_tweet_child_tweet FOREIGN KEY (child_tweet_id) REFERENCES tweet(id)
);

CREATE TABLE user_session (
  id SERIAL PRIMARY KEY,
  user_id INTEGER UNIQUE NOT NULL,
  token TEXT UNIQUE NOT NULL,
  expire_date TIMESTAMP NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_user_session_user FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE retweet (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  tweet_id INTEGER NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT unique_retweet UNIQUE (user_id, tweet_id),
  CONSTRAINT fk_retweet_user FOREIGN KEY (user_id) REFERENCES user(id),
  CONSTRAINT fk_retweet_tweet FOREIGN KEY (tweet_id) REFERENCES tweet(id)
);

CREATE TABLE likes (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  tweet_id INTEGER NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT unique_likes UNIQUE (user_id, tweet_id),
  CONSTRAINT fk_likes_user FOREIGN KEY (user_id) REFERENCES user(id),
  CONSTRAINT fk_likes_tweet FOREIGN KEY (tweet_id) REFERENCES tweet(id)
);
