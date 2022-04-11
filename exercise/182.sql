select
    distinct p1.Email
from
    Person p1,
    Person p2
where
    p1.Email = p2.Email
    and
    p1.Id != p2.Id

/* 这个快 */
select 
    email 
from 
    person 
group by 
    email 
having 
    count(email)>1