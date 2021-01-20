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
    - `-c` - umożliwia określenie maksymalnych odległości między miastami 0..n i generację specjalnego przypadku, w którym miasta leżą blisko siebie (domyślnie wyłączone)

	- `-g` - wybor analizowanego algorytmu (domyslnie 0):
		- `0` - algorytm liniowy
		- `1` - algorytm brutalny
		- `2` - algorytm brutalny wspolbiezny
- W trybie `m3` parametry trybu `m2` i dodatkowo:
	- `-k` - liczba roznych wielkosci `n` w sekwencji generowanych problemow (domyslnie 1)
	- `-t` - roznica miedzy kolejnymi rozmiarami problemu `n` w sekwencji (step) (domyslnie 500)
	- `-r` - liczba wygenerowanych instacji problemu dla kazdego `n` (domyslnie 1)

Przyklad uruchomienia: `./AAL_Project -m3 -n1000 -k30 -t500 -r10 -w10000`

## Zastosowane algorytmy
### Algorytm brutalny
Dla każdego z n miast koszt jest liczony osobno - da to rozwiązanie o złożoności kwadratowej, ponieważ obliczenie kosztu dla danego miasta będzie miało złożoność liniową, a samych miast jest liniowo wiele. Rozwiązanie to jest wolniejsze, jednak jest prostsze w implementacji. Dzięki temu będzie można sprawdzić poprawność rozwiązań o mniejszej złożoności obliczeniowej. Złożoność pamięciowa tego algorytmu jest liniowa, ponieważ korzysta on z tablic wejściowych, które są liniowo duże. Zmiennych pomocniczych wykorzystywanych w trakcie algorytmu (np. przechowujących najniższy koszt i indeks miasta, który ten koszt wygenerowało) jest stała liczba, zatem złożoność pamięciowa jest liniowa.
W ramach stworzonego programu powstał także współbieżny algorytm brutalny. Idea jego działania jest bardzo podobna do zwykłego algorytmu brutalnego. Jedyną różnicą jest fakt, że algorytm współbieżny dzieli zbiór miast na k grup, gdzie k to liczba partycji, i dla każdej grupy tworzy nowy wątek (zrealizowany za pomocą goroutines). Dla każdego z miast z danej grupy wywoływana jest w nowo utworzonym wątku funkcją liniowo wyliczająca koszt dla danego miasta. Dzięki zastosowaniu współbieżności przy 4 partycjach możemy liczyć na około 2-krotnie krótszy czas działania algorytmu.

### Algorytm liniowy
W ramach preprocessingu zostanie przygotowana tablica sum prefiksowych odległości między miastami, tak aby w czasie stałym odpowiadać na pytanie o odległość między dowolnymi dwoma miastami. Najpierw dla jednego miasta, oznaczonego indeksem 0, wyliczony w czasie liniowym zostanie sumaryczny koszt wysyłki wszystkich ciężarówek wyjeżdżających z tego miasta “w lewo” (L) i “w prawo” (P). Następnie wykonane zostanie n-1 kroków przejścia z miasta i do i+1, w każdym kroku w zamortyzowanym czasie stałym zostaną wyznaczone wartości L i P dla miasta i+1, na podstawie wartości L i P dla miasta i.
- di-(i+1) - odległość między miastami i oraz i+1
- zi - zapotrzebowanie miasta i
- L - koszt transportu ładunku wysłanego po okręgu zgodnie z ruchem wskazówek zegara
- P - koszt transportu ładunku wysłanego po okręgu przeciwnie do ruchu wskazówek zegara
- S - grupa miast w zaznaczonym wycinku okręgu - miasta, do których kierunek transportu zmieni się w danym kroku


Krok przejścia z miasta i do i+1 na przykładzie przejścia z miasta 0 do miasta 1, gdzie znamy wartości L i P dla miasta 0 i obliczamy je (L’ i P’) dla miasta 1:
- Nowy koszt L’ i P’:
    - L’ = L - d0-1*(z1+z2+z3+z4) + z5*d1-5 + z6*d1-6 + z7*d1-7
    - P’ = P + d0-1*(z8+z9+z0) - z5*d0-5 - z6*d0-6 - z7*d0-7

## Pliki źródłowe
Projekt został zdekomponowany na kilka plików źródłowych:
- `main.go` - główny kod programu, parsowanie opcji programu, uruchomienie odpowiednich algorytmów i podsumowanie rezultatów
- `bruteForce.go` - algorytm brutalny
- `concurrentBruteForce.go` - algorytm brutalny zaimplementowany współbieżnie z użyciem wbudowanych w go funkcji
- `linear.go` - algorytm optymalny
- `ui.go` - wyświetlenie opcji programu i wczytanie danych w trybie 1


## Informacje dodatkowe
W trakcie pomiaru czasu algorytmu optymalnego dla małych n (n <= 1000) dodatkowo wykonywany jest algorytm brutalny, aby potwierdzić poprawność algorytmu optymalnego. Algorytm brutalny wykonywany jest po zakończeniu pomiaru czasu działania algorytmu optymalnego.

Dane testowe były generowane losowo w zakresach określonych przez parametry uruchomienia. Wygenerowane zostały również dane testowe obejmujące przypadki szczególne - miasta oddalone od siebie o niewielkie odległości w jednej części okręgu.