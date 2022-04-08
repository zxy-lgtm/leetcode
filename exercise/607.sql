select
    s.name
from
    salesperson s
where
    s.sales_id not in
    (select sales_id from orders o,company c where c.name = 'RED' and c.com_id = o.com_id)
