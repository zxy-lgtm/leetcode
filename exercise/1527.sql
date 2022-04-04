select 
    patient_id,
    patient_name,
    conditions
from
    Patients
where
    conditions like '% DIAB1%'
    or
    conditions like 'DIAB1%'

/*正则慢了一点*/

select 
    patient_id,
    patient_name,
    conditions
from
    Patients
where
    conditions rlike '^DIAB1|.*\\sDIAB1' 

