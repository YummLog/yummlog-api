create table foodposts
(
    id              text not null
        constraint foodposts_pk
            primary key,
    restaurant_name text not null,
    address1        text,
    address2        text,
    city            text,
    state           text,
    country         text,
    zipcode         text,
    user_id         text,
    created_by      text,
    created_date    date,
    updated_by      text,
    updated_date    date,
    notes           text
);




create table postdetails
(
    id         text not null
        constraint postdetails_pk
            primary key,
    post_id    text not null
        constraint postdetails_foodposts_id_fk
            references foodposts,
    item       text not null,
    experience text not null
);





