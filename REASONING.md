## 1. Project Setup and Learning Curve

I started this project quite late and had to familiarize myself with `sqlc` first. Initially, it was confusing because I expected it to behave like a typical ORM, but it actually works in the opposite way. It generates Go methods from raw SQL queries, rather than generating SQL from Go structs.

I decided to make `sqlc` use the same schema as my migrations, since that felt intuitive and avoided duplicating the schema.

I was also not familiar with Fiber, but the challenging part was adhering to `repository-service-handler` structure provided.

---

## 2. Project Structure Decisions

The given project structure felt a bit overkill for the scope. I was confused why the `db/` package was outside `internal/`.
I had to adhere to it so, I organized it like so:

* `db/sqlc` & `db/migrations`: Contains SQLC queries and migrations
* `internal/database/`: COntains DB connection and sqlc-generated code

This keeps the low-level database code internal and prevents it from being imported directly elsewhere.

---

## 3. Models

I chose to create my own response models like `UserCreateResponse` and `UserUpdateResponse`, since, you need to include `age` in the respective endpoints

---

## 4. Repository Layer & Service Layer

First, I thought there was no need for repository layer, since `sqlc` essentially acts as one. Then, I realized that `sqlc` returns errors as is. So, I decided to use the repo layer as a **wrapper** over `sqlc`. I used this wrapper to normalize errors, like using `ErrUserNotFound`. Id later implement errors like: `ErrUserAlreadyExists`, `ErrConflict`, etc

The **Service Layer** is where I implemented the logic to calculating a user’s age from their DOB.

---

## 6. Handlers

HTTP Layer is where I parse the incoming JSON, validate it, handle any errors, and then delegate the remaining stuff to service layer.

I also used a custom Date type which is basically time.Time
```go
type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error 

func (d Date) MarshalJSON() ([]byte, error)
```

TLDR: This allows the API to accept `YYYY-MM-DD` string and return error for invalid or empty dates, while keeping the internal representation as `time.Time`.

**Longer explanantion**:

When Go is decoding a particual type to JSON, it sees if the type has a `UnmarshalJSON` method on it. If it does, Go will call this method to determine how to decode the value.

So, I created a `Date` type, and implemented my own `UnmarshalJSON()` on it, which Go uses to decode JSON into this type(Date). 
The same concept applied to `MarshalJSON`