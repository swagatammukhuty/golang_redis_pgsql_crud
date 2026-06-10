# Redis Caching Strategy Comparison

This project demonstrates a CRUD application using **Go**, **Fiber**, **PostgreSQL**, and **Redis**. The primary database is PostgreSQL, while Redis is used as a caching layer to improve read performance.

## Dataset Used

To test the caching strategy, 30,000 users were inserted into PostgreSQL using:

```sql
INSERT INTO users (id, name, email)
SELECT
    i,
    'User' || i,
    'user' || i || '@example.com'
FROM generate_series(1, 30000) AS s(i);
```

---

# Approach 1: Single Redis Key for All Users

Commit:

`28da1092d1afb5facc45c381e6dd103fd5ed95d1`

In this implementation, all user records are stored in a single Redis key:

```text
all_users
```

Example:

```json
[
  {
    "ID": 1,
    "name": "User1",
    "email": "user1@example.com"
  },
  {
    "ID": 2,
    "name": "User2",
    "email": "user2@example.com"
  }
]
```

### Request Flow

#### Get All Users

```text
GET /getAllUsers
    ↓
Redis GET all_users
    ↓
Return cached data
```

This approach works well when the application frequently requests the complete user list.

### Limitation

Suppose we need to fetch:

```text
User29999
```

Since Redis stores all users in a single JSON array, the application must:

```text
Redis GET all_users
    ↓
Deserialize all 30,000 users
    ↓
Iterate through the collection
    ↓
Find User29999
```

Even though Redis retrieval is fast, the application still performs unnecessary deserialization and iteration over thousands of records to locate a single user.

As the dataset grows, this becomes less efficient.

---

# Approach 2: One Redis Key Per User

Commit:
Last One

In this implementation, each user is stored in Redis using a dedicated key.

Example:

```text
user:1
user:2
user:3
...
user:29999
user:30000
```

Each key contains only the corresponding user's data:

```json
{
  "ID": 29999,
  "name": "User29999",
  "email": "user29999@example.com"
}
```

### Request Flow

#### Get Individual User

```text
GET /getUser?id=29999
    ↓
Redis GET user:29999
    ↓
Return user data
```

No iteration or full dataset deserialization is required.

### Benefits

* Direct access to individual users.
* Constant-time key lookup.
* Reduced memory transfer between Redis and application.
* Better scalability as the number of users grows.
* Closer to real-world production caching strategies.

---

# Comparison

| Feature             | Single Key (`all_users`)        | Per User Key (`user:<id>`)      |
| ------------------- | ------------------------------- | ------------------------------- |
| Get all users       | Fast                            | Requires aggregation            |
| Get individual user | Requires iteration              | Direct lookup                   |
| Memory transfer     | High                            | Low                             |
| Scalability         | Moderate                        | Better                          |
| Production usage    | Small datasets / simple caching | Common for object-level caching |

---

# Conclusion

The single-key approach is useful for learning Redis concepts and works well for small datasets or when the application primarily retrieves all users at once.

However, for applications that frequently access individual users, storing each user under a dedicated Redis key is more efficient and scalable. This approach avoids deserializing large datasets and provides direct access to user records, which is why it is commonly used in production environments.
