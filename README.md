# Orchard Odyssey

```sql
INSERT INTO accounts (username, email, password_hash) VALUES ('jacques', 'gkounkou@gmail.com', 'b47fddd607ca735ad44ef235637efa15ed76f94d3760d0ee33b7c50f70957886')

INSERT INTO regions (region_name) VALUES
    ('Afghanistan'),
    ('Albania'),
    ('Algeria'),
    ('Andorra'),
    ('Angola'),
    ('Antigua and Barbuda'),
    ('Argentina'),
    ('Armenia'),
    ('Australia'),
    ('Austria'),
    ('Azerbaijan'),
    ('Bahamas'),
    ('Bahrain'),
    ('Bangladesh'),
    ('Barbados'),
    ('Belarus'),
    ('Belgium'),
    ('Belize'),
    ('Benin'),
    ('Bhutan'),
    ('Bolivia'),
    ('Bosnia and Herzegovina'),
    ('Botswana'),
    ('Brazil'),
    ('Brunei'),
    ('Bulgaria'),
    ('Burkina Faso'),
    ('Burundi'),
    ('Cabo Verde'),
    ('Cambodia'),
    ('Cameroon'),
    ('Canada'),
    ('Central African Republic'),
    ('Chad'),
    ('Chile'),
    ('China'),
    ('Colombia'),
    ('Comoros'),
    ('Congo, Democratic Republic of the'),
    ('Congo, Republic of the'),
    ('Costa Rica'),
    ('Croatia'),
    ('Cuba'),
    ('Cyprus'),
    ('Czech Republic'),
    ('Denmark'),
    ('Djibouti'),
    ('Dominica'),
    ('Dominican Republic'),
    ('East Timor'),
    ('Ecuador'),
    ('Egypt'),
    ('El Salvador'),
    ('Equatorial Guinea'),
    ('Eritrea'),
    ('Estonia'),
    ('Eswatini'),
    ('Ethiopia'),
    ('Fiji'),
    ('Finland'),
    ('France'),
    ('Gabon'),
    ('Gambia'),
    ('Georgia'),
    ('Germany'),
    ('Ghana'),
    ('Greece'),
    ('Grenada'),
    ('Guatemala'),
    ('Guinea'),
    ('Guinea-Bissau'),
    ('Guyana'),
    ('Haiti'),
    ('Honduras'),
    ('Hungary'),
    ('Iceland'),
    ('India'),
    ('Indonesia'),
    ('Iran'),
    ('Iraq'),
    ('Ireland'),
    ('Israel'),
    ('Italy'),
    ('Jamaica'),
    ('Japan'),
    ('Jordan'),
    ('Kazakhstan'),
    ('Kenya'),
    ('Kiribati'),
    ('Korea, North'),
    ('Korea, South'),
    ('Kosovo'),
    ('Kuwait'),
    ('Kyrgyzstan'),
    ('Laos'),
    ('Latvia'),
    ('Lebanon'),
    ('Lesotho'),
    ('Liberia'),
    ('Libya'),
    ('Liechtenstein'),
    ('Lithuania'),
    ('Luxembourg'),
    ('Madagascar'),
    ('Malawi'),
    ('Malaysia'),
    ('Maldives'),
    ('Mali'),
    ('Malta'),
    ('Marshall Islands'),
    ('Mauritania'),
    ('Mauritius'),
    ('Mexico'),
    ('Micronesia'),
    ('Moldova'),
    ('Monaco'),
    ('Mongolia'),
    ('Montenegro'),
    ('Morocco'),
    ('Mozambique'),
    ('Myanmar'),
    ('Namibia'),
    ('Nauru'),
    ('Nepal'),
    ('Netherlands'),
    ('New Zealand'),
    ('Nicaragua'),
    ('Niger'),
    ('Nigeria'),
    ('North Macedonia'),
    ('Norway'),
    ('Oman'),
    ('Pakistan'),
    ('Palau'),
    ('Palestine'),
    ('Panama'),
    ('Papua New Guinea'),
    ('Paraguay'),
    ('Peru'),
    ('Philippines'),
    ('Poland'),
    ('Portugal'),
    ('Qatar'),
    ('Romania'),
    ('Russia'),
    ('Rwanda'),
    ('Saint Kitts and Nevis'),
    ('Saint Lucia'),
    ('Saint Vincent and the Grenadines'),
    ('Samoa'),
    ('San Marino'),
    ('Sao Tome and Principe'),
    ('Saudi Arabia'),
    ('Senegal'),
    ('Serbia'),
    ('Seychelles'),
    ('Sierra Leone'),
    ('Singapore'),
    ('Slovakia'),
    ('Slovenia'),
    ('Solomon Islands'),
    ('Somalia'),
    ('South Africa'),
    ('South Sudan'),
    ('Spain'),
    ('Sri Lanka'),
    ('Sudan'),
    ('Suriname'),
    ('Sweden'),
    ('Switzerland'),
    ('Syria'),
    ('Taiwan'),
    ('Tajikistan'),
    ('Tanzania'),
    ('Thailand'),
    ('Togo'),
    ('Tonga'),
    ('Trinidad and Tobago'),
    ('Tunisia'),
    ('Turkey'),
    ('Turkmenistan'),
    ('Tuvalu'),
    ('Uganda'),
    ('Ukraine'),
    ('United Arab Emirates'),
    ('United Kingdom'),
    ('USA'),
    ('Uruguay'),
    ('Uzbekistan'),
    ('Vanuatu'),
    ('Vatican City'),
    ('Venezuela'),
    ('Vietnam'),
    ('Yemen'),
    ('Zambia'),
    ('Zimbabwe');

INSERT INTO fruit_vegetables (fruit_vegetable_name, category, region_name, image_url) VALUES
    ('Apple', 'fruit', 'USA', ''),
    ('Banana', 'fruit', 'Brazil', ''),
    ('Orange', 'fruit', 'Spain', ''),
    ('Dragon fruit', 'fruit', 'Vietnam', ''),
    ('Mango', 'fruit', 'India', ''),
    ('Pineapple', 'fruit', 'Philippines', ''),
    ('Grapes', 'fruit', 'Italy', ''),
    ('Strawberry', 'fruit', 'France', ''),
    ('Blueberry', 'fruit', 'Canada', ''),
    ('Cherry', 'fruit', 'Turkey', ''),
    ('Pear', 'fruit', 'China', ''),
    ('Peach', 'fruit', 'Japan', ''),
    ('Plum', 'fruit', 'Australia', ''),
    ('Papaya', 'fruit', 'South Africa', ''),
    ('Avocado', 'fruit', 'Mexico', ''),
    ('Kiwi', 'fruit', 'New Zealand', ''),
    ('Watermelon', 'fruit', 'Egypt', ''),
    ('Cantaloupe', 'fruit', 'Morocco', ''),
    ('Lemon', 'fruit', 'Greece', ''),
    ('Lime', 'fruit', 'Peru', ''),
    ('Apricot', 'fruit', 'Turkey', ''),
    ('Fig', 'fruit', 'Greece', ''),
    ('Pomegranate', 'fruit', 'India', ''),
    ('Raspberry', 'fruit', 'France', ''),
    ('Blackberry', 'fruit', 'Canada', ''),
    ('Grapefruit', 'fruit', 'USA', ''),
    ('Kumquat', 'fruit', 'China', ''),
    ('Guava', 'fruit', 'India', ''),
    ('Passion fruit', 'fruit', 'Brazil', ''),
    ('Date', 'fruit', 'Egypt', ''),
    ('Persimmon', 'fruit', 'China', ''),
    ('Coconut', 'fruit', 'Indonesia', ''),
    ('Jackfruit', 'fruit', 'India', ''),
    ('Durian', 'fruit', 'Thailand', ''),
    ('Rhubarb', 'vegetable', 'Canada', ''),
    ('Carrot', 'vegetable', 'USA', ''),
    ('Spinach', 'vegetable', 'China', ''),
    ('Lettuce', 'vegetable', 'Italy', ''),
    ('Broccoli', 'vegetable', 'Italy', ''),
    ('Cauliflower', 'vegetable', 'USA', ''),
    ('Cucumber', 'vegetable', 'India', ''),
    ('Tomato', 'vegetable', 'Italy', ''),
    ('Bell pepper', 'vegetable', 'Mexico', ''),
    ('Zucchini', 'vegetable', 'Italy', ''),
    ('Pumpkin', 'vegetable', 'Mexico', ''),
    ('Squash', 'vegetable', 'USA', ''),
    ('Beetroot', 'vegetable', 'USA', ''),
    ('Sweet corn', 'vegetable', 'USA', ''),
    ('Onion', 'vegetable', 'China', ''),
    ('Garlic', 'vegetable', 'China', ''),
    ('Leek', 'vegetable', 'Morocco', ''),
    ('Asparagus', 'vegetable', 'Argentina', ''),
    ('Artichoke', 'vegetable', 'Italy', ''),
    ('Brussels sprouts', 'vegetable', 'Belgium', ''),
    ('Chili pepper', 'vegetable', 'Mexico', ''),
    ('Radish', 'vegetable', 'USA', ''),
    ('Ackee', 'fruit', 'Jamaica', ''),
    ('Miracle Fruit', 'fruit', 'Ghana', ''),
    ('Horned Melon', 'fruit', 'Kenya', ''),
    ('Mangosteen', 'fruit', 'Thailand', ''),
    ('Rambutan', 'fruit', 'Malaysia', ''),
    ('Salak', 'fruit', 'Indonesia', ''),
    ('Romanesco', 'vegetable', 'Italy', ''),
    ('Oca', 'vegetable', 'Peru', ''),
    ('Sunchoke', 'vegetable', 'USA', ''),
    ('Chayote', 'vegetable', 'Mexico', '');

INSERT INTO account_fruit_vegetables (account_name, fruit_vegetable_name) VALUES
		('jacques', 'Kumquat'),
        ('jacques', 'Persimmon'),
        ('jacques', 'Durian')

INSERT INTO information (information_name, fruit_vegetable_name, description) VALUES
    ('Apple','Apple','A sweet, edible fruit produced by an apple tree.'),
    ('Banana','Banana','A long, curved fruit with a yellow skin.'),
    ('Orange','Orange','A citrus fruit with a tough, bright orange skin.'),
    ('Dragon fruit','Dragon fruit','A tropical fruit known for its unique appearance and taste.'),
    ('Mango','Mango','A juicy, sweet fruit with a tropical flavor.'),
    ('Pineapple','Pineapple','A tropical fruit with a spiky skin and sweet flesh.'),
    ('Grapes','Grapes','Small, round fruits that can be red, green, or purple.'),
    ('Strawberry','Strawberry','A red, juicy fruit with a sweet flavor.'),
    ('Blueberry','Blueberry','A small, round fruit with a blue color and sweet taste.'),
    ('Cherry','Cherry','A small, round fruit with a sweet taste, usually red or black.'),
    ('Pear','Pear','A sweet fruit with a rounded base and tapering top.'),
    ('Peach','Peach','A juicy fruit with a soft skin and sweet flavor.'),
    ('Plum','Plum','A round fruit with a smooth skin and sweet flavor.'),
    ('Papaya','Papaya','A tropical fruit with a sweet, musky flavor.'),
    ('Avocado','Avocado','A creamy fruit with a rich flavor, often used in salads.'),
    ('Kiwi','Kiwi','A small, brown fruit with a green interior and unique flavor.'),
    ('Watermelon','Watermelon','A large, juicy fruit with a green rind and red flesh.'),
    ('Cantaloupe','Cantaloupe','A type of melon with a sweet, orange flesh.'),
    ('Lemon','Lemon','A sour, yellow citrus fruit.'),
    ('Lime','Lime','A small, green citrus fruit with a tart flavor.'),
    ('Apricot','Apricot','A small, orange fruit with a sweet taste.'),
    ('Fig','Fig','A sweet fruit with a unique texture and flavor.'),
    ('Pomegranate','Pomegranate','A fruit with a tough outer skin and juicy seeds inside.'),
    ('Raspberry','Raspberry','A small, red fruit with a tart flavor.'),
    ('Blackberry','Blackberry','A dark purple fruit with a sweet and tart taste.'),
    ('Grapefruit','Grapefruit','A citrus fruit with a tangy flavor and pink flesh.'),
    ('Kumquat','Kumquat','A small, orange citrus fruit that is eaten whole.'),
    ('Guava','Guava','A tropical fruit with a sweet flavor and fragrant aroma.'),
    ('Passion fruit','Passion fruit','A tropical fruit with a sweet and tangy taste.'),
    ('Date','Date','A sweet fruit that comes from the date palm tree.'),
    ('Persimmon','Persimmon','A sweet fruit with a smooth texture and vibrant color.'),
    ('Coconut','Coconut','A tropical fruit with a hard shell and sweet, milky interior.'),
    ('Jackfruit','Jackfruit','A large fruit with a spiky outer skin and sweet, fleshy interior.'),
    ('Durian','Durian','A tropical fruit known for its strong smell and rich, creamy texture.'),
    ('Ackee', 'Ackee', 'A tropical fruit that is a key ingredient in the national dish of Jamaica, ackee and saltfish. It must be prepared carefully, as parts of the fruit are toxic if not properly cooked.'),
    ('Miracle Fruit', 'Miracle Fruit', 'A small red berry that causes sour foods to taste sweet after being consumed, due to a protein called miraculin.'),
    ('Horned Melon', 'Horned Melon', 'A spiky, orange fruit with a jelly-like, lime-green interior. Also known as kiwano, it has a mildly sweet and tart flavor.'),
    ('Mangosteen', 'Mangosteen', 'A tropical fruit with a sweet, tangy flavor and a thick, inedible rind. Known as the "queen of fruits", it is highly prized in Southeast Asia.'),
    ('Rambutan', 'Rambutan', 'A hairy, red fruit with a sweet and juicy white interior. It is related to the lychee and is popular in Southeast Asia.'),
    ('Salak', 'Salak', 'Also known as snake fruit due to its reddish-brown scaly skin, this fruit has a sweet and tangy flavor, and a crunchy, juicy texture.'),
    ('Romanesco', 'Romanesco', 'A chartreuse-colored vegetable with a fractal pattern. It is a variant of cauliflower and has a nutty, slightly crunchy texture.'),
    ('Oca', 'Oca', 'A tuber that is a staple in Andean cuisine. It is brightly colored and has a tangy flavor, often eaten boiled, baked, or fried.'),
    ('Sunchoke', 'Sunchoke', 'Also known as Jerusalem artichoke, this tuber has a sweet, nutty flavor and can be eaten raw or cooked. It is native to North America.'),
    ('Chayote', 'Chayote', 'A green, pear-shaped squash that is a staple in Latin American cuisine. It has a mild flavor and is often cooked like a vegetable.'),
    ('Rhubarb', 'Rhubarb', 'A vegetable with tart, pink stalks often used in pies.'),
    ('Carrot', 'Carrot', 'A root vegetable known for its orange color and sweet taste.'),
    ('Spinach', 'Spinach', 'A leafy green vegetable rich in vitamins and minerals.'),
    ('Lettuce', 'Lettuce', 'A leafy green vegetable used in salads and sandwiches.'),
    ('Broccoli','Broccoli','A green vegetable with a tree-like structure and mild flavor.'),
    ('Cauliflower','Cauliflower','A white vegetable with a mild flavor and versatile uses.'),
    ('Cucumber','Cucumber','A crisp, green vegetable often used in salads and pickles.'),
    ('Tomato','Tomato','A red fruit often used as a vegetable in salads and cooking.'),
    ('Bell pepper','Bell pepper','A colorful vegetable with a sweet flavor, available in various colors.'),
    ('Zucchini','Zucchini','A green summer squash with a mild flavor.'),
    ('Pumpkin','Pumpkin','A large, orange vegetable with a sweet and nutty flavor.'),
    ('Squash','Squash','A versatile vegetable available in various types and flavors.'),
    ('Beetroot','Beetroot','A root vegetable with a deep red color and earthy taste.'),
    ('Sweet corn','Sweet corn','A sweet and juicy vegetable often eaten on the cob.'),
    ('Onion','Onion','A pungent vegetable used to flavor a variety of dishes.'),
    ('Garlic','Garlic','A strong-smelling vegetable used as a seasoning and flavoring.'),
    ('Leek','Leek','A vegetable with a mild onion flavor, often used in soups.'),
    ('Asparagus','Asparagus','A green vegetable with tender stalks and a distinct flavor.'),
    ('Artichoke','Artichoke','A vegetable with edible buds and a rich, nutty flavor.'),
    ('Brussels sprouts','Brussels sprouts','Small, green vegetables with a slightly bitter taste.'),
    ('Chili pepper','Chili pepper','A hot vegetable used to add heat and flavor to dishes.'),
    ('Radish','Radish','A crunchy, spicy vegetable often used in salads.');


curl -X POST http://localhost:8000/create-account- -d "username=isabelle&email=isabelle@gmail.com&password_hash=067642f78c613c7c1aece9cabffff8215c92bca1bb82578595b1806fd0ff20ca"
curl -X DELETE "http://localhost:8000/delete-account-?account_id=isabelle"
curl -X POST http://localhost:8000/notification-stats-b47fddd607ca735ad44ef235637efa15ed76f94d3760d0ee33b7c50f70957886
curl -X POST http://localhost:8000/add-unknown-items-b47fddd607ca735ad44ef235637efa15ed76f94d3760d0ee33b7c50f70957886 -d "account_id=jacques" -d "fruit_vegetable_name=DragonFruit" -d "fruit_vegetable_name=Kale"
curl -X DELETE "http://localhost:8000/delete-unknown-items-b47fddd607ca735ad44ef235637efa15ed76f94d3760d0ee33b7c50f70957886?account_id=jacques&fruit_vegetable_name=DragonFruit"


curl -X POST http://localhost:8000/connect -d "username=sofia&hash=067642f78c613c7c1aece9cabffff8215c92bca1bb82578595b1806fd0ff20ca"
curl -X GET http://localhost:8000/notification-suggestion-067642f78c613c7c1aece9cabffff8215c92bca1bb82578595b1806fd0ff20ca -b cookies.txt
curl -X POST http://localhost:8000/disconnect -b cookies.txt
curl -X GET http://localhost:8000/notification-suggestion-067642f78c613c7c1aece9cabffff8215c92bca1bb82578595b1806fd0ff20ca -b cookies.txt
```