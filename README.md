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
```
W przypadku przyjętej przez nas numeracji miast od 0

## TODO
- generowanie danych losowe, ale dodac tez przypadki szczegolne i na nich tez sprawdzic

- czysty kod i czysta kompilacja
- dokumentacja: Wraz z programem źródłowym trzeba przekazać dokumentację końcową w formie elektronicznej (dokument Word,OpenOfficelub pdf)  zawierającą  opis  problemu  i  metody (lub metod) rozwiązania, opis wykorzystywanych struktur danych i algorytmów pomocniczych oraz ocenę  spodziewanej  złożoności  algorytmu(wykorzystaną  w  trybie  wykonania  3.).

## Add to readme
- "wizytówka" studenta
- krótka specyfikacjaproblemu (albo przynajmniej tytuł problemu)
- informacja o możliwych poleceniach aktywacji programu (opcje, parametry, ...)
- opis konwencji dotyczących danych wejściowych i prezentacji wynikówkrótki opis metody rozwiązania,  zastosowanych algorytmów i struktur danych
- informacje  o  funkcjonalnej  dekompozycji  programu  na  moduły  źródłowe -nagłówkowe  i implementacyjne ("przewodnik" po plikach źródłowych)
- inne informacje dodatkowe o szczególnych decyzjach projektowych (np. ograniczenia dotyczące rozmiaru  problemu,  charakterystyki  generatorów  danych  testowych,  specjalne  konwencje  w alokacji  pamięci  dynamicznej,  wymagania  dotyczące  typów  parametryzujących    szablony,  konwencje związane z obsługą sytuacji wyjątkowych, itp.)


## DONE
- UI z wyswietlaniem pomocy
- 3 tryby wywolania: 
    - full - z wejscia pobierane wszystko, mozna np przekierowac plik na wejscie
    - random - okreslenie metaparametrow przez uzytkownika
    - presentation - program samodzielnie uruchamia random dla roznych n wybranych przez uzytkownika (np opisane jako wartosc minimalna, liczba testow i step miedzy n-ami, dla kazdego n-a 10 uruchomien), stworzenie tabeli wynikow oceny zlozonosci. W tabeli 3 kolumny: ni (n posortowane rosnaco od gory na dol), czas wykonania t(ni), q(n) = t(n)/(c*T(n)) gdzie c = t(n_mediana)/T(n_mediana)
- sprawdzanie poprawnosci brutem w solveForAnalysis przy malych n (zeby automatycznie sprawdzic liniowke)
- wybor algo do analizy w CLI
