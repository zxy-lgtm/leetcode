select 
    name
from
    customer
where
    referee_id != 2
    or
    referee_id is null
/*！=会自动去掉null值的那一列*/