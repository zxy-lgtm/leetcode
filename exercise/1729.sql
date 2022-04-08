select
    user_id,
    count(distinct follower_id) followers_count
from
    followers
group by
    user_id