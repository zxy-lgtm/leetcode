select
    name,
    ifnull(sum(distance),0) travelled_distance
from
    users
left join
    rides
on
    users.id = rides.user_id
group by
    users.id
order by
    travelled_distance desc,
    name 