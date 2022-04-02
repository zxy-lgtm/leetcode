select
    name Customers
from 
    Customers
where Id not in
    (select CustomerId from Orders);

/*not in 效率较低*/

select c.Name as Customers 
from Customers as c
left join Orders as o on c.Id = o.CustomerId
where o.Id is null
/* 效率较高 */
