import requests

def get_location(ip_address):
    # ipstack API anahtarýnýzý buraya ekleyin
    api_key = 'APÝ anahtarýnýzý girmelisiniz'
    url = f'http://api.ipstack.com/{ip_address}?access_key={api_key}'

    response = requests.get(url)
    data = response.json()

    if response.status_code == 200:
        return {
            'IP': data.get('ip'),
            'Country': data.get('country_name'),
            'Region': data.get('region_name'),
            'City': data.get('city'),
            'Latitude': data.get('latitude'),
            'Longitude': data.get('longitude')
        }
    else:
        return {'Error': data.get('error', {}).get('info', 'An error occurred')}

# Kullanýcýdan IP adresi alýnýr
ip_address = input("Lütfen konumunu öðrenmek istediðiniz IP adresini girin: ")
location = get_location(ip_address)

if 'Error' in location:
    print(f"Bir hata oluþtu: {location['Error']}")
else:
    print(f"IP Adresi: {location['IP']}")
    print(f"Ülke: {location['Country']}")
    print(f"Bölge: {location['Region']}")
    print(f"Þehir: {location['City']}")
    print(f"Enlem: {location['Latitude']}")
    print(f"Boylam: {location['Longitude']}")
