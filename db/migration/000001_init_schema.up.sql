Table users{
  id      int [pk, increment]
  name    varchar(50)
}

Table follow{
  id          int [pk, increment]
  user_id     int 
  follow_id   int
}

Ref: users.id < follow.user_id
