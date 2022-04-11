select 
    actor_id,
    director_id
from
    actordirector
group by
    actor_id,
    director_id
having
    count(director_id)>2