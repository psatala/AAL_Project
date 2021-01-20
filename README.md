# Epidemia - Projekt AAL
## Piotr Libera, Piotr Satała

Władze pewnej wyspy na Pacyfiku zdecydowały się położyć kres epidemii koronawirusa poprzez znaczące zwiększenie dostępności płynów do dezynfekcji (do stosowania wewnętrznego i zewnętrznego). W tym celu zdecydowano o szybkim wybudowaniu fabryki płynów dezynfekcyjnych,
z której każdego dnia ciężarówki będą je rozwozić do miast znajdujących się na wyspie. Wyspa zamieszkana jest jedynie w pasie przybrzeżnym, zaś wszystkie miasta połączone są pojedynczą drogą, tworzącą okrąg. Wiadomo, jakie będzie dzienne zapotrzebowanie każdego
miasta na płyn, zaś władze muszą zdecydować jaka powinna być lokalizacja fabryki, tak aby zminimalizować koszty transportu.

Dane są - lista D zawierająca liczby naturalne określające odległości pomiędzy kolejnymi miastami w kilometrach i tym samym definiująca ich położenia na drodze przybrzeżnej (oraz określająca całkowitą długość tej drogi), oraz lista Z zawierająca także liczby naturalne,
określające dzienne zapotrzebowanie poszczególnych miast na płyn dezynfekcyjny (jednostką jest tu jedna ciężarówka). Koszt transportu jednej ciężarówki na odległość 1km jest stały. Miast jest nie więcej niż 10000, zapotrzebowanie każdego z nich nie przekracza
1000 ciężarówek dziennie, zaś całkowita długość drogi nie jest większa niż 1000000 km. Zaproponuj algorytm podający miasto w którym należy wybudować fabrykę, tak aby dzienny koszt transportu był minimalny (oraz obliczający ten koszt).

Przykładowe dane wejściowe
```
3
1 4 7
10 13 5
```
Wyjście
```
1
30
```
W przypadku przyjętej przez nas numeracji miast od 0


## Uruchomienie programu
`./AAl_Project [Tryb uruchomienia] [Opcje trybu uruchomienia]`

`[tryb uruchomienia]` moze przyjac wartosc:
- `-m1` - wszystkie dane pobierane ze standardowego wejscia
- `-m2` - pobierane parametry problemu i generacja tablic D i Z
- `-m3` - tryb prezentacji, w ktorym program bada zlozonosc algorytmu generujac sekwencje rozwazanych wielkosci problemu n. Program tworzy tabelę wyników oceny złożoności. W tabeli są 3 kolumny:
    - ni (n posortowane rosnąco)
    - czas wykonania t(ni)
    - `q(n) = t(n)/(c*T(n))`, gdzie `c = t(n_mediana)/T(n_mediana)`

Opcje uruchomienia:
- W trybie `m1` nie ma dodatkowych opcji, na standardowym wejsciu powinny pojawic sie liczby oznaczajace kolejno:
	- `n`
	- `n` liczb odpowiadajacych odleglosciom `D` miedzy kolejnymi miastami, zaczynajac od odleglosci miedzy miastami 0 i 1
	- `n` liczb odpowiadajacych zapotrzebowanie `Z` kolejnych miast od 0 do n-1
- W trybie `m2`:
	- `-n` - wielkosc problemu (domyslnie 1000)
	- `-w` - maksymalna wartosc liczb w `D` i `Z` (domyslnie 10000)
	- `-s` - wartosc seed dla generatora liczb pseudolosowych (domyslnie aktualny czas systemowy)
	- `-g` - wybor analizowanego algorytmu (domyslnie 0):
		- `0` - algorytm liniowy
		- `1` - algorytm brutalny
		- `2` - algorytm brutalny wspolbiezny
- W trybie `m3` parametry trybu `m2` i dodatkowo:
	- `-k` - liczba roznych wielkosci `n` w sekwencji generowanych problemow (domyslnie 1)
	- `-t` - roznica miedzy kolejnymi rozmiarami problemu `n` w sekwencji (step) (domyslnie 500)
	- `-r` - liczba wygenerowanych instacji problemu dla kazdego `n` (domyslnie 1)

Przyklad uruchomienia: `./AAL_Project -m3 -n1000 -k30 -t500 -r10 -w10000`

## TODO Zastosowane algorytmy i struktury danych 
krótki opis metody rozwiązania,  zastosowanych algorytmów i struktur danych

## Pliki źródłowe
Projekt został zdekomponowany na kilka plików źródłowych:
- `main.go` - główny kod programu, parsowanie opcji programu, uruchomienie odpowiednich algorytmów i podsumowanie rezultatów
- `bruteForce.go` - algorytm brutalny
- `concurrentBruteForce.go` - algorytm brutalny zaimplementowany współbieżnie z użyciem wbudowanych w go funkcji
- `linear.go` - algorytm optymalny
- `ui.go` - wyświetlenie opcji programu i wczytanie danych w trybie 1


## TODO Informacje dodatkowe
W trakcie pomiaru czasu algorytmu optymalnego dla małych n (n <= 1000) dodatkowo wykonywany jest algorytm brutalny, aby potwierdzić poprawność algorytmu optymalnego. Algorytm brutalny wykonywany jest po zakończeniu pomiaru czasu działania algorytmu optymalnego.


TODO inne informacje dodatkowe o szczególnych decyzjach projektowych (np. ograniczenia dotyczące rozmiaru  problemu,  charakterystyki  generatorów  danych  testowych,  specjalne  konwencje  w alokacji  pamięci  dynamicznej,  wymagania  dotyczące  typów  parametryzujących    szablony,  konwencje związane z obsługą sytuacji wyjątkowych, itp.)




## TODO
- generowanie danych losowe, ale dodac tez przypadki szczegolne i na nich tez sprawdzic
