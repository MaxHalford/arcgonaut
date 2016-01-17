import pandas as pd

airportNames = ['id', 'name', 'city', 'country', 'IATA', 'ICAO',
                'lat', 'lon', 'altitude', 'timezone', 'DST', 'tz']
routeNames = ['airline', 'airlineID', 'source', 'sourceID',
              'destination', 'destinationID', 'codeshare', 'stops',
              'equipment']
base = 'https://raw.githubusercontent.com/jpatokal/openflights/master/data/'
airports = pd.read_csv(base + 'airports.dat', names=airportNames)
routes = pd.read_csv(base + 'routes.dat', names=routeNames)
outwards = pd.merge(left=pd.merge(left=airports, right=routes, left_on='IATA',
                    right_on='source'), right=airports, left_on='destination',
                    right_on='IATA')[['country_x', 'country_y']]
outwards.columns = ('source', 'destination')
outwards['both'] = outwards['source'] + '_' + outwards['destination']
with open('airports.arcgo', 'w') as arcgo:
    for countries, count in outwards.groupby('both').count().iterrows():
        a, b = countries.split('_')
        amount = count['destination']
        if a != b:
            arcgo.write('{0}>{1}>{2}\n'.format(a, b, amount))
