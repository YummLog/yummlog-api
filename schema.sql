create table foodposts
(
    id               text not null
        constraint foodposts_pk
            primary key,
    "restaurantName" text not null,
    address1         text,
    address2         text,
    city             text,
    state            text,
    country          text,
    zipcode          text,
    "userId"         text,
    "createdBy"      text,
    "createdDate"    date,
    "updatedBy"      text,
    "updatedDate"    date,
    notes            text
);


create table postdetails
(
    id         text not null
        constraint postdetails_pk
            primary key,
    "postId"   text not null
        constraint postdetails_foodposts_id_fk
            references foodposts,
    item       text not null,
    experience text not null
);



