alter table horo_service
drop column meeting_location ;

alter table horo_service
add column meeting_location geometry;