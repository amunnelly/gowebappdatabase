select team as Team,
sum(Points) as Points,
sum(GoalDifference) as GD,
sum(GoalsFor) as GoalsFor,
sum(GoalsAgainst) as GoalsAgainst
from points
where season = '%s'
and Competition = 'Premier League'
and Country = 'England'
group by team
order by points desc;
