# Instrukcje obsługi aplikacji do wystawiania faktur dla portalu fakturownia.pl

## Wprowadzenie

Aplikacja umożliwiająca wystawianie faktur dla portalu fakturownia.pl jest prostym narzędziem do tworzenia faktur oraz zarządzania nimi przy użyciu API udostępnionego przez portal fakturownia.pl. Poniżej znajdują się instrukcje, które pomogą Ci rozpocząć korzystanie z aplikacji.

## Konfiguracja

Aby korzystać z aplikacji, musisz skonfigurować dostęp do API fakturownia.pl. Oto jak to zrobić:

1. Zaloguj się do swojego konta na portalu fakturownia.pl.
2. Przejdź do ustawień konta lub profilu.
3. Znajdź sekcję "API" lub "Klucze API".
4. Wygeneruj nowy klucz API (jeśli go jeszcze nie masz).
5. Skopiuj wygenerowany klucz API. 
6. Skopiuj twoje id Usera, przykładowo: https://agent89.fakturownia.pl <- w tym przypadku `domain` to agent89.
7. Wejdź w zakładke klienci, https://agent89.fakturownia.pl/clients i wybierz klienta dla którego chcesz wygenerować faktury.
8. Po wybraniu klienta trzeba skopiować jego `client_id` https://agent89.fakturownia.pl/clients/88732555 <- w tym przypadku jest to  `88732555`.


Teraz, gdy masz `api_key`, `domain`, `client_id` możesz go użyć w aplikacji.

## Uruchamianie aplikacji

Teraz, gdy masz wszystkie wymagane informacje: `api_key`, `domain`, `client_id`, możesz je wykorzystać w aplikacji.
Aby uruchomić aplikację i wystawić fakturę:

1. **Pobierz** program wraz z całą zawartością.
2. Przejdź do pliku `data.json`.
3. **Uzupełnij** plik `data.json` danymi z wcześniejszej konfiguracji (`api_key`, `domain`, `client_id` itp.).
4. **Uzupełnij** plik `invoices.csv`. Przykład znajduje się w pliku `example.csv`.
5. **Uruchom** program `bombelaio-fakturownia.exe`.
6. **Korzystaj** z wygenerowanych faktur.

---

> **Informacja**  
> Upewnij się, że wprowadzane dane są prawidłowe i aktualne, aby wystawiane faktury były poprawnie generowane.
