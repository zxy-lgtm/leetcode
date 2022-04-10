select
    user_id buyer_id,
    join_date,
    count(order_id) orders_in_2019
from
    users
left join
    orders
on
    orders.buyer_id = users.user_id
and
    year(order_date)='2019'
group by
    user_id