select
    name,
    sum(amount) balance
from
    users,
    transactions
where
    users.account = transactions.account
group by
    users.account
having
    balance > 10000