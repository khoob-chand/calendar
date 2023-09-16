 create table calender(
    start_date timestamp,
    end_date timestamp,
     day text,
     mode text,
     occurence int,
     selectcase text,
     title text,
     week text);
 ---- create above / drop below ----
 drop table calender;
 ---- create above / drop below ----
 -- Write your migrate down statements here. If this migration is irreversible
 -- Then delete the separator line above.