select  Processors.id, Processors.name from Personal
left join Abteilungen on Abteilungen.id = Personal.abteilungId
left join Items on Items.depId = Abteilungen.id
right join Processors on Processors.id = Items.processorId
where Processors.id in (
	select  Processors.id from  Items
	right join Processors on Processors.id = Items.processorId
    where Processors.enabled=1
    and Items.enabled=1    
	group by Processors.id
) and (
	Abteilungen.id = Personal.abteilungId
    or abteilungId = 1
)
group by Processors.id 