select
    player_id,
    min(event_date) first_login
from
    activity
group by
    player_id

/* min还可以比较日期啊神奇 */