package seedgenerator

var (
	words = [...]string{
		"абак",
		"абзац",
		"аборт",
		"абсолютный",
		"абстрактный",
		"абсурд",
		"август",
		"авиационный",
		"автор",
		"агат",
		"агент",
		"агрессор",
		"адекватный",
		"адрес",
		"ажур",
		"азия",
		"азот",
		"аист",
		"академия",
		"аккуратный",
		"актер",
		"активный",
		"актуальный",
		"акция",
		"албания",
		"алкоголь",
		"алтарь",
		"алый",
		"альтернативный",
		"альянс",
		"алюминиевый",
		"амбар",
		"американец",
		"ампула",
		"амулет",
		"анализ",
		"ангар",
		"ангел",
		"ангина",
		"английский",
		"анонс",
		"апельсин",
		"апендицит",
		"апостол",
		"аппарат",
		"апрель",
		"аптека",
		"арабский",
		"арбитраж",
		"арбуз",
		"арест",
		"аркан",
		"армия",
		"аромат",
		"артист",
		"арфа",
		"архитектура",
		"аршин",
		"асбест",
		"аспирант",
		"асфальт",
		"атака",
		"атаман",
		"атеист",
		"атлас",
		"атлет",
		"атмосфера",
		"атомный",
		"атрибут",
		"аттестат",
		"аудит",
		"аукцион",
		"афиша",
		"афоризм",
		"африка",
		"ацетон",
		"аэрофлот",
		"баба",
		"бабушка",
		"багаж",
		"багет",
		"багор",
		"базар",
		"базовый",
		"баланс",
		"балет",
		"балл",
		"банан",
		"бандит",
		"банк",
		"бант",
		"баня",
		"бард",
		"баржа",
		"барон",
		"барс",
		"батон",
		"баян",
		"бдительный",
		"бегать",
		"бегун",
		"беда",
		"бедный",
		"бедро",
		"бежать",
		"бездна",
		"безопасный",
		"безумный",
		"бекон",
		"белила",
		"белка",
		"белый",
		"берег",
		"беркут",
		"беседа",
		"бесконечный",
		"бесплатный",
		"бессмысленный",
		"бетон",
		"бешеный",
		"библиотека",
		"бидон",
		"бизнес",
		"бизон",
		"билет",
		"биография",
		"биология",
		"биржа",
		"бирка",
		"бисер",
		"бить",
		"благодарный",
		"блаженный",
		"бланк",
		"бледный",
		"блестящий",
		"ближний",
		"близкий",
		"блин",
		"блок",
		"бобр",
		"богатый",
		"бодрость",
		"боевой",
		"боец",
		"божественный",
		"божий",
		"бокал",
		"боковой",
		"бокс",
		"болеть",
		"болт",
		"большой",
		"бомба",
		"бордюр",
		"борзая",
		"бормотать",
		"бороться",
		"борт",
		"борщ",
		"борьба",
		"босой",
		"ботаник",
		"ботинок",
		"боярин",
		"бояться",
		"брак",
		"браслет",
		"брат",
		"бревно",
		"бред",
		"брелок",
		"бригада",
		"бриз",
		"британский",
		"брови",
		"бродить",
		"бром",
		"бронзовый",
		"бросать",
		"брусок",
		"брызги",
		"брюки",
		"брюхо",
		"бубен",
		"бублик",
		"бугры",
		"будка",
		"будни",
		"будущий",
		"буйвол",
		"буква",
		"букет",
		"буклет",
		"буксир",
		"булат",
		"булка",
		"бумага",
		"бунт",
		"бурный",
		"буря",
		"бусы",
		"бутылка",
		"бухта",
		"бушлат",
		"бывать",
		"бывший",
		"былой",
		"быстрый",
		"бытовой",
		"быть",
		"бюджет",
		"бюрократ",
		"бюст",
		"вагон",
		"важный",
		"ваза",
		"вазелин",
		"валенок",
		"вальс",
		"валютный",
		"валяться",
		"ваниль",
		"ванна",
		"варвар",
		"вариант",
		"вата",
		"ваучер",
		"вафля",
		"вахта",
		"вброс",
		"ввести",
		"вводить",
		"ведро",
		"ведущий",
		"везти",
		"велеть",
		"величина",
		"велосипед",
		"венок",
		"вера",
		"верба",
		"верить",
		"верный",
		"вероятный",
		"версия",
		"вертикальный",
		"верхний",
		"вершина",
		"веселый",
		"весна",
		"вести",
		"весы",
		"ветер",
		"ветхий",
		"вечер",
		"вечный",
		"вещество",
		"вещь",
		"взаимный",
		"взвод",
		"взгляд",
		"вздохнуть",
		"взнос",
		"взрослый",
		"взрыв",
		"взять",
		"видеть",
		"видимый",
		"видный",
		"визит",
		"вилла",
		"вино",
		"вираж",
		"вирус",
		"висеть",
		"виски",
		"висок",
		"витамин",
		"виток",
		"витрина",
		"вихрь",
		"вишня",
		"вклад",
		"включать",
		"вкус",
		"владеть",
		"влажный",
		"власть",
		"влиять",
		"вложить",
		"вмятина",
		"внедрение",
		"внезапный",
		"внести",
		"внешний",
		"внимание",
		"вносить",
		"внук",
		"внутренний",
		"вода",
		"водитель",
		"водка",
		"водный",
		"воевать",
		"военный",
		"возвращение",
		"возглавлять",
		"воздух",
		"возможный",
		"возникать",
		"возраст",
		"воинский",
		"война",
		"войско",
		"войти",
		"волна",
		"волос",
		"волшебный",
		"вольный",
		"воля",
		"вооруженный",
		"вопрос",
		"ворота",
		"ворс",
		"воскликнуть",
		"воспитание",
		"восстановить",
		"восток",
		"вояж",
		"впадать",
		"впечатление",
		"впрок",
		"враг",
		"врать",
		"врач",
		"вредный",
		"время",
		"всадник",
		"всевозможный",
		"вселенная",
		"всемирный",
		"всеобщий",
		"всероссийский",
		"вскочить",
		"вспомнить",
		"встать",
		"встреча",
		"вступать",
		"всходы",
		"всякое",
		"всяческий",
		"втиснуть",
		"вторник",
		"втулка",
		"вулкан",
		"вход",
		"вцепиться",
		"вчера",
		"выбирать",
		"выбор",
		"выбрать",
		"вывести",
		"вывод",
		"выглядеть",
		"выгодный",
		"выдать",
		"выделить",
		"выжить",
		"вызвать",
		"вызов",
		"вызывать",
		"выиграть",
		"выйти",
		"вынести",
		"выносить",
		"вынудить",
		"выпасть",
		"выпить",
		"выполнять",
		"выпуск",
		"выразить",
		"вырваться",
		"высказать",
		"выслушать",
		"высота",
		"выставка",
		"высший",
		"вытащить",
		"вытянуть",
		"выход",
		"выявить",
		"выяснить",
		"вязать",
		"вяленый",
		"гавань",
		"гадюка",
		"газета",
		"газовый",
		"гараж",
		"гвоздь",
		"гейзер",
		"генерал",
		"гениальный",
		"географический",
		"геологический",
		"герб",
		"германский",
		"герой",
		"гетры",
		"гетто",
		"гибкий",
		"гигантский",
		"гимнаст",
		"гипс",
		"гирлянда",
		"гиря",
		"гитара",
		"глава",
		"гладкий",
		"глаз",
		"глина",
		"глобальный",
		"глубина",
		"глупый",
		"глухой",
		"глыба",
		"глядеть",
		"глянуть",
		"гнать",
		"гнездо",
		"гном",
		"говорить",
		"годовой",
		"голова",
		"голубой",
		"голый",
		"гонка",
		"гора",
		"горб",
		"гордый",
		"гореть",
		"горло",
		"гормон",
		"горный",
		"город",
		"горький",
		"горячий",
		"господин",
		"гость",
		"государство",
		"готовый",
		"гражданин",
		"грамотный",
		"граница",
		"грех",
		"греческий",
		"грозить",
		"громкий",
		"грош",
		"грубый",
		"грудь",
		"грузовой",
		"",
		"группа",
		"грустный",
		"грядущий",
		"грязный",
		"губа",
		"губернатор",
		"гулять",
		"гуманитарный",
		"густой",
		"давать",
		"давление",
		"давний",
		"далекий",
		"дальний",
		"дама",
		"данные",
		"дать",
		"дача",
		"дверь",
		"двигатель",
		"движение",
		"двинуться",
		"двойной",
		"двор",
		"девочка",
		"девушка",
		"дежурный",
		"действие",
		"декабрь",
		"декоративный",
		"делать",
		"делиться",
		"дело",
		"демонстрация",
		"денежный",
		"день",
		"депутат",
		"дерево",
		"держать",
		"десяток",
		"деталь",
		"детский",
		"дешевый",
		"деятельность",
		"диван",
		"дикий",
		"дипломатический",
		"директор",
		"длина",
		"длиться",
		"дневной",
		"добавить",
		"добиться",
		"добро",
		"доверять",
		"довольный",
		"догадаться",
		"договор",
		"дождь",
		"дожидаться",
		"дойти",
		"доказать",
		"доклад",
		"доктор",
		"документ",
		"долг",
		"должен",
		"доллар",
		"доложить",
		"доля",
		"домашний",
		"дополнительный",
		"допустить",
		"дорога",
		"доска",
		"достать",
		"доход",
		"дочка",
		"дочь",
		"драгоценный",
		"драматический",
		"древний",
		"дрожать",
		"друг",
		"дружба",
		"думать",
		"дурак",
		"дурной",
		"духовный",
		"душа",
		"душевный",
		"дыхание",
		"дышать",
		"дядя",
		"еврей",
		"европейский",
		"единица",
		"ежегодный",
		"ежедневный",
		"ездить",
		"естественный",
		"есть",
		"ехать",
		"жалеть",
		"жалкий",
		"жаловаться",
		"жаркий",
		"ждать",
		"желать",
		"железный",
		"желтый",
		"жена",
		"жениться",
		"женский",
		"женщина",
		"жертва",
		"жесткий",
		"живот",
		"жидкий",
		"жизнь",
		"жилищный",
		"жилой",
		"жилье",
		"жирный",
		"житель",
		"жить",
		"журнал",
		"жуткий",
		"забавный",
		"забота",
		"забрать",
		"забыть",
		"завести",
		"зависеть",
		"зависимость",
		"завод",
		"загадочный",
		"заглянуть",
		"заговорить",
		"задача",
		"задержать",
		"задний",
		"задуматься",
		"зайти",
		"заказ",
		"заключение",
		"закон",
		"закрыть",
		"закурить",
		"заложить",
		"заметить",
		"замок",
		"занимать",
		"занятие",
		"запад",
		"запись",
		"заплатить",
		"запомнить",
		"запретить",
		"заработать",
		"зарегистрировать",
		"зарплата",
		"зарубежный",
		"заседание",
		"заслужить",
		"засмеяться",
		"заснуть",
		"заставить",
		"засыпать",
		"затрата",
		"захватить",
		"захотеть",
		"защита",
		"заявление",
		"звать",
		"звезда",
		"звонок",
		"звук",
		"звучать",
		"здание",
		"здешний",
		"здоровье",
		"здравый",
		"зеленый",
		"земельный",
		"земля",
		"земной",
		"зенитный",
		"зеркало",
		"зима",
		"зимний",
		"злой",
		"знак",
		"знаменитый",
		"знание",
		"знать",
		"значение",
		"золото",
		"зона",
		"зрение",
		"зритель",
		"игра",
		"идеальный",
		"идеологический",
		"идея",
		"идти",
		"избавиться",
		"избежать",
		"избирательный",
		"избрать",
		"известие",
		"извинить",
		"издание",
		"излишний",
		"изложить",
		"изменить",
		"изображение",
		"израильский",
		"изучать",
		"изящный",
		"иметь",
		"импортный",
		"имущество",
		"инвестиционный",
		"индийский",
		"инженер",
		"инициатива",
		"иностранный",
		"институт",
		"интеллектуальный",
		"интерес",
		"интимный",
		"информация",
		"иркутский",
		"искать",
		"исключение",
		"искренний",
		"искусство",
		"испанский",
		"исполнение",
		"испугаться",
		"испытание",
		"исследовать",
		"истинный",
		"история",
		"исходный",
		"исчезать",
		"итальянский",
		"итог",
		"июль",
		"июнь",
		"кабинет",
		"кавказский",
		"кадр",
		"казаться",
		"казенный",
		"камень",
		"канал",
		"кандидат",
		"капитан",
		"карман",
		"карта",
		"касаться",
		"категория",
		"качество",
		"квадратный",
		"квартира",
		"кивать",
		"кивнуть",
		"киевский",
		"километр",
		"кино",
		"кинуться",
		"кирпичный",
		"китайский",
		"класс",
		"клетка",
		"клиент",
		"клинический",
		"клуб",
		"ключ",
		"книга",
		"князь",
		"кодекс",
		"кожа",
		"колесо",
		"количество",
		"коллектив",
		"колоссальный",
		"кольцо",
		"колючий",
		"командир",
		"комитет",
		"коммерческий",
		"комната",
		"компания",
		"комплекс",
		"комсомольский",
		"конец",
		"конкурс",
		"конструкция",
		"контакт",
		"конфликт",
		"концерт",
		"кончить",
		"корабль",
		"корень",
		"коридор",
		"кормить",
		"король",
		"корпус",
		"космический",
		"коснуться",
		"кость",
		"кофе",
		"край",
		"красный",
		"краткий",
		"кредит",
		"кремлевский",
		"крепкий",
		"крест",
		"кривой",
		"кризис",
		"крик",
		"криминальный",
		"критический",
		"кричать",
		"кровь",
		"крохотный",
		"крошечный",
		"круг",
		"крупный",
		"крутой",
		"крыло",
		"крыша",
		"кулак",
		"культура",
		"купить",
		"курить",
		"курс",
		"кусок",
		"куст",
		"кухня",
		"лагерь",
		"ладонь",
		"ласковый",
		"латинский",
		"левый",
		"легендарный",
		"легкий",
		"ледяной",
		"лежать",
		"лезть",
		"лейтенант",
		"лекарственный",
		"лесной",
		"лестница",
		"летать",
		"лететь",
		"летний",
		"лето",
		"лечить",
		"лечь",
		"либеральный",
		"лидер",
		"линейный",
		"линия",
		"лист",
		"литература",
		"лицо",
		"личность",
		"лишить",
		"лишний",
		"ловить",
		"логический",
		"лодка",
		"ложиться",
		"ложный",
		"локальный",
		"ломать",
		"лошадь",
		"лунный",
		"лучший",
		"лысый",
		"любить",
		"любовь",
		"людской",
		"люсин",
		"магазин",
		"магнитный",
		"майор",
		"максимальный",
		"маленький",
		"малый",
		"мальчик",
		"мама",
		"мамин",
		"март",
		"масло",
		"масса",
		"мастер",
		"масштаб",
		"материал",
		"мать",
		"махнуть",
		"машина",
		"мгновение",
		"медицинский",
		"медленный",
		"медный",
		"международный",
		"мелкий",
		"меньший",
		"менять",
		"мера",
		"мероприятие",
		"мертвый",
		"место",
		"месяц",
		"металл",
		"метод",
		"метр",
		"механизм",
		"мечта",
		"мешать",
		"мешок",
		"милиция",
		"миллион",
		"милый",
		"минеральный",
		"министр",
		"миновать",
		"минута",
		"мирный",
		"мировой",
		"младший",
		"мнение",
		"многолетний",
		"множество",
		"мобильный",
		"могучий",
		"модель",
		"модный",
		"мозг",
		"мокрый",
		"молиться",
		"молоко",
		"молчать",
		"момент",
		"монастырь",
		"моральный",
		"море",
		"морской",
		"москвич",
		"мост",
		"мочь",
		"мощность",
		"мрачный",
		"мудрый",
		"мужик",
		"мужской",
		"мужчина",
		"музей",
		"музыка",
		"муниципальный",
		"мутный",
		"мучить",
		"мысль",
		"мыть",
		"мягкий",
		"мясной",
		"мясо",
		"набирать",
		"наблюдение",
		"набор",
		"набрать",
		"надежда",
		"надоесть",
		"название",
		"назначение",
		"называть",
		"наибольший",
		"наивный",
		"найти",
		"наличие",
		"налог",
		"намерен",
		"нанести",
		"напечатать",
		"написать",
		"напомнить",
		"направо",
		"нарисовать",
		"народ",
		"нарушение",
		"население",
		"настроение",
		"натуральный",
		"наука",
		"научный",
		"находить",
		"национальный",
		"начало",
		"начальник",
		"начинать",
		"небесный",
		"небо",
		"неверный",
		"невидимый",
		"невозможный",
		"невысокий",
		"негативный",
		"недавний",
		"неделя",
		"недостаток",
		"нежный",
		"независимый",
		"незнакомый",
		"неизвестный",
		"нелепый",
		"немалый",
		"немец",
		"немыслимый",
		"ненавидеть",
		"ненужный",
		"необычный",
		"неожиданный",
		"неплохой",
		"непонятный",
		"непростой",
		"нервный",
		"нести",
		"несчастный",
		"неудачный",
		"нефть",
		"нехороший",
		"неясный",
		"нижний",
		"низкий",
		"ничтожный",
		"новенький",
		"новость",
		"новый",
		"нога",
		"номер",
		"норма",
		"носить",
		"ночной",
		"ночь",
		"ноябрь",
		"нравиться",
		"нуждаться",
		"нужный",
		"нынешний",
		"обед",
		"обернуться",
		"обеспечить",
		"обещать",
		"обидеть",
		"область",
		"обмануть",
		"обнаружить",
		"обнять",
		"обозначить",
		"обойти",
		"оборона",
		"образ",
		"обрести",
		"обстановка",
		"обсудить",
		"обусловить",
		"обучение",
		"обширный",
		"общаться",
		"общество",
		"общий",
		"объединение",
		"объект",
		"объявление",
		"обыкновенный",
		"обычный",
		"обязанность",
		"оглянуться",
		"огненный",
		"огонь",
		"ограничить",
		"огромный",
		"одежда",
		"одетый",
		"одинаковый",
		"ожидание",
		"озеро",
		"означать",
		"оказать",
		"окно",
		"окончание",
		"округ",
		"октябрь",
		"олимпийский",
		"опасность",
		"операция",
		"опираться",
		"описание",
		"оплата",
		"определение",
		"оптимальный",
		"опубликовать",
		"опустить",
		"опыт",
		"оранжевый",
		"орать",
		"орган",
		"оригинальный",
		"оружие",
		"осветить",
		"освободить",
		"осень",
		"основа",
		"особый",
		"остаток",
		"осторожный",
		"остров",
		"осуществить",
		"ответ",
		"отдать",
		"отдел",
		"отдохнуть",
		"отдых",
		"отец",
		"отечественный",
		"отказ",
		"открытие",
		"отличие",
		"отложить",
		"отмена",
		"отнести",
		"отношение",
		"отобрать",
		"отозваться",
		"отойти",
		"оторвать",
		"отправить",
		"отпуск",
		"отрасль",
		"отрезать",
		"отрицательный",
		"отсутствие",
		"отчаянный",
		"офицер",
		"оформить",
		"охватить",
		"охрана",
		"оценка",
		"очевидный",
		"очередь",
		"очки",
		"ошибка",
		"ощутить",
		"ощущение",
		"падать",
		"пакет",
		"палата",
		"палец",
		"память",
		"папа",
		"парад",
		"парень",
		"парижский",
		"парк",
		"парламентский",
		"партия",
		"пахнуть",
		"педагогический",
		"пенсия",
		"первый",
		"перевод",
		"период",
		"персональный",
		"песня",
		"песок",
		"пестрый",
		"петербургский",
		"петь",
		"печать",
		"пиво",
		"писатель",
		"пистолет",
		"письмо",
		"питерский",
		"пить",
		"пищевой",
		"плавать",
		"плакать",
		"план",
		"пластиковый",
		"плата",
		"плечо",
		"плоский",
		"плотный",
		"плохой",
		"площадь",
		"плыть",
		"победа",
		"побывать",
		"поверхность",
		"повод",
		"повседневный",
		"повторять",
		"повышение",
		"погибнуть",
		"поглядеть",
		"поговорить",
		"пограничный",
		"подарок",
		"подводный",
		"подготовка",
		"поддержка",
		"поделиться",
		"подземный",
		"подлинный",
		"подмосковный",
		"поднять",
		"подойти",
		"подписать",
		"подруга",
		"подтвердить",
		"подумать",
		"подход",
		"подчеркнуть",
		"подъезд",
		"поезд",
		"поехать",
		"пожать",
		"пожелать",
		"пожилой",
		"позвать",
		"поздний",
		"позиция",
		"познакомиться",
		"поинтересоваться",
		"поиск",
		"поймать",
		"пойти",
		"показатель",
		"покинуть",
		"покой",
		"покрыть",
		"покупатель",
		"полагать",
		"поле",
		"политика",
		"полковник",
		"полный",
		"полоса",
		"получка",
		"польза",
		"полюбить",
		"помещение",
		"помнить",
		"помощь",
		"понадобиться",
		"понести",
		"понимать",
		"понравиться",
		"понятие",
		"пообещать",
		"попадать",
		"попросить",
		"популярный",
		"попытка",
		"пора",
		"порог",
		"портрет",
		"поручить",
		"порядок",
		"посадить",
		"посвятить",
		"поселок",
		"посидеть",
		"последствие",
		"посмотреть",
		"посоветовать",
		"пост",
		"посылать",
		"потеря",
		"поток",
		"потребность",
		"потянуть",
		"похожий",
		"поцеловать",
		"почва",
		"почетный",
		"почитать",
		"почтовый",
		"почувствовать",
		"поэзия",
		"поэт",
		"появиться",
		"пояснить",
		"правда",
		"праздник",
		"практика",
		"превратить",
		"предмет",
		"прежний",
		"президент",
		"прекрасный",
		"премия",
		"преодолеть",
		"препарат",
		"прервать",
		"преступление",
		"прибыть",
		"привычный",
		"приглашение",
		"придумать",
		"прием",
		"прижать",
		"признак",
		"приказ",
		"приличный",
		"пример",
		"принцип",
		"приобрести",
		"природа",
		"приступ",
		"приходить",
		"причина",
		"приятель",
		"проблема",
		"провод",
		"программа",
		"продукт",
		"проект",
		"прожить",
		"прозрачный",
		"производство",
		"пройти",
		"прокурор",
		"промышленность",
		"пропасть",
		"просьба",
		"противник",
		"профессия",
		"прохлада",
		"процент",
		"прочный",
		"прошлое",
		"прощать",
		"проявить",
		"прыгать",
		"прямой",
		"прятать",
		"психология",
		"птица",
		"публика",
		"пугать",
		"пункт",
		"пускать",
		"пустыня",
		"путь",
		"пушкинский",
		"пыль",
		"пытаться",
		"пышный",
		"пьеса",
		"пьяный",
		"работа",
		"равный",
		"радикальный",
		"радость",
		"разбить",
		"развитие",
		"разговор",
		"разделить",
		"различный",
		"размер",
		"разница",
		"разобрать",
		"разработка",
		"разумный",
		"район",
		"ракета",
		"рамка",
		"ранить",
		"ранний",
		"раскрыть",
		"рассказ",
		"растение",
		"расход",
		"расчет",
		"расширить",
		"рациональный",
		"рваться",
		"реагировать",
		"реакция",
		"реальность",
		"ребенок",
		"ребята",
		"революция",
		"регион",
		"регулярный",
		"редактор",
		"редкий",
		"режим",
		"резиновый",
		"резкий",
		"результат",
		"река",
		"реклама",
		"рекомендовать",
		"религиозный",
		"ремонт",
		"республика",
		"ресторан",
		"ресурс",
		"реформа",
		"речь",
		"решение",
		"римский",
		"риск",
		"рисовать",
		"рисунок",
		"ровный",
		"родина",
		"родной",
		"родственник",
		"рождение",
		"розовый",
		"роковой",
		"роль",
		"роман",
		"роскошный",
		"российский",
		"рост",
		"рубеж",
		"рубль",
		"рука",
		"руководить",
		"русский",
		"рухнуть",
		"ручка",
		"ручной",
		"рыба",
		"рыбный",
		"рыжий",
		"рынок",
		"рядовой",
		"садиться",
		"садовый",
		"самарский",
		"самолет",
		"сапог",
		"саратовский",
		"сбить",
		"сбор",
		"сбросить",
		"сведение",
		"свежий",
		"свернуть",
		"свести",
		"свет",
		"свидетель",
		"свобода",
		"своеобразный",
		"свойство",
		"связь",
		"святой",
		"священный",
		"сдавать",
		"сдать",
		"сделка",
		"северный",
		"сегодняшний",
		"седой",
		"сезон",
		"секрет",
		"секс",
		"секунда",
		"село",
		"сельский",
		"семья",
		"сентябрь",
		"сердце",
		"середина",
		"серия",
		"серый",
		"сестра",
		"сеть",
		"сибирский",
		"сигарета",
		"сигнал",
		"сидеть",
		"сила",
		"силовой",
		"сильный",
		"симпатичный",
		"синий",
		"система",
		"ситуация",
		"сиять",
		"сказка",
		"складываться",
		"склонный",
		"скорость",
		"скромный",
		"скрытый",
		"скучный",
		"слабый",
		"слава",
		"сладкий",
		"след",
		"слеза",
		"слепой",
		"слово",
		"сложный",
		"слой",
		"сломать",
		"служба",
		"слух",
		"случай",
		"слушать",
		"слышать",
		"смелый",
		"смена",
		"смерть",
		"смех",
		"смотреть",
		"смочь",
		"смутный",
		"смысл",
		"снег",
		"снежный",
		"снижение",
		"снимать",
		"сниться",
		"снять",
		"собака",
		"собирать",
		"собрание",
		"собственность",
		"событие",
		"совесть",
		"совместный",
		"совпадать",
		"современный",
		"соглашение",
		"содержание",
		"соединение",
		"сожаление",
		"создание",
		"сознание",
		"сойти",
		"сократить",
		"солдат",
		"соленый",
		"солидный",
		"солнце",
		"сомнение",
		"сонный",
		"сообщение",
		"соответствие",
		"сопровождать",
		"сосед",
		"состав",
		"сотня",
		"сотрудник",
		"сохранить",
		"социальный",
		"сочинение",
		"союз",
		"спасать",
		"спать",
		"спектакль",
		"специалист",
		"спешить",
		"спина",
		"список",
		"сплошной",
		"спокойный",
		"спор",
		"способ",
		"справа",
		"спросить",
		"спрятать",
		"спуск",
		"сравнение",
		"среда",
		"срок",
		"срочный",
		"стабильный",
		"ставка",
		"стакан",
		"стальной",
		"станция",
		"старик",
		"статья",
		"стекло",
		"стена",
		"степень",
		"стесняться",
		"стиль",
		"стихи",
		"стоимость",
		"стол",
		"сторона",
		"стоять",
		"страна",
		"стрелок",
		"строка",
		"структура",
		"студент",
		"стул",
		"стучать",
		"субъект",
		"судебный",
		"судить",
		"судьба",
		"сума",
		"суметь",
		"сумка",
		"сумма",
		"сунуть",
		"суровый",
		"сутки",
		"суть",
		"сухой",
		"существо",
		"сущность",
		"сфера",
		"сформулировать",
		"схватить",
		"схема",
		"сход",
		"сцена",
		"счастье",
		"счесть",
		"счет",
		"считать",
		"съесть",
		"сыграть",
		"сырой",
		"сытый",
		"сюжет",
		"таблица",
		"таинственный",
		"тайна",
		"талант",
		"таможенный",
		"танец",
		"танк",
		"танцевать",
		"татарский",
		"тащить",
		"твердый",
		"творчество",
		"театр",
		"текст",
		"текущий",
		"телефон",
		"тело",
		"тема",
		"темнота",
		"температура",
		"тенденция",
		"тень",
		"теория",
		"теплый",
		"терпеть",
		"территория",
		"терять",
		"тесный",
		"тетя",
		"техника",
		"течение",
		"течь",
		"типичный",
		"тихий",
		"тишина",
		"товар",
		"тогдашний",
		"толпа",
		"толстый",
		"тоненький",
		"тонкий",
		"торговля",
		"торжественный",
		"торопиться",
		"торчать",
		"точка",
		"точный",
		"трава",
		"трагический",
		"традиция",
		"транспорт",
		"тратить",
		"требование",
		"тревожный",
		"трезвый",
		"трогать",
		"тронуть",
		"труба",
		"труд",
		"тупой",
		"турецкий",
		"тысяча",
		"тюремный",
		"тюрьма",
		"тяжелый",
		"тяжкий",
		"тянуть",
		"убедить",
		"убежать",
		"убивать",
		"убийство",
		"убить",
		"убрать",
		"уважать",
		"увеличение",
		"уверять",
		"увидеть",
		"угол",
		"угроза",
		"удаваться",
		"удар",
		"удаться",
		"удачный",
		"удержаться",
		"удивление",
		"удобный",
		"удовольствие",
		"уезжать",
		"уехать",
		"ужас",
		"узкий",
		"узнавать",
		"уйти",
		"указ",
		"украсть",
		"улица",
		"уличный",
		"улыбка",
		"умереть",
		"уметь",
		"умирать",
		"умный",
		"умственный",
		"университет",
		"уникальный",
		"уничтожить",
		"упасть",
		"упомянуть",
		"управление",
		"уровень",
		"урок",
		"усилие",
		"условие",
		"услуга",
		"услышать",
		"усмехнуться",
		"уснуть",
		"успех",
		"успокоить",
		"усталый",
		"устойчивый",
		"устроить",
		"уступать",
		"утвердить",
		"уточнить",
		"утратить",
		"утренний",
		"утро",
		"уход",
		"участок",
		"учебный",
		"ученик",
		"учесть",
		"учет",
		"учитель",
		"учреждение",
		"уютный",
		"факт",
		"фамилия",
		"фантастический",
		"февраль",
		"федерация",
		"фигура",
		"физический",
		"философия",
		"фильм",
		"финансовый",
		"финский",
		"фирма",
		"фонд",
		"форма",
		"фотография",
		"фраза",
		"французский",
		"фронт",
		"фундаментальный",
		"функция",
		"футбол",
		"характер",
		"хват",
		"хвост",
		"химический",
		"хитрый",
		"хлеб",
		"ходить",
		"хозяин",
		"холодный",
		"хороший",
		"хотеть",
		"хохотать",
		"храм",
		"хранить",
		"христианский",
		"хрупкий",
		"художник",
		"худший",
		"царский",
		"царь",
		"цвет",
		"целевой",
		"целое",
		"целый",
		"цель",
		"цена",
		"ценить",
		"ценность",
		"центр",
		"церковь",
		"цивилизованный",
		"цифра",
		"чайный",
		"часть",
		"часы",
		"человек",
		"человечество",
		"чемпионат",
		"черный",
		"черта",
		"честь",
		"четкий",
		"чеченский",
		"чиновник",
		"число",
		"чистый",
		"читатель",
		"член",
		"чрезвычайный",
		"чтение",
		"чувство",
		"чудесный",
		"чудо",
		"чуждый",
		"чужой",
		"шанс",
		"шахматы",
		"шелковый",
		"шептать",
		"широкий",
		"школа",
		"штаб",
		"штат",
		"штука",
		"шумный",
		"шутить",
		"шутка",
		"щека",
		"экологический",
		"экономика",
		"экран",
		"эксперт",
		"электрон",
		"элемент",
		"эмоциональный",
		"энергия",
		"эпоха",
		"эстетический",
		"этаж",
		"этап",
		"этнический",
		"эффект",
		"южный",
		"юный",
		"юридический",
		"яблоко",
		"явиться",
		"явление",
		"являться",
		"явный",
		"ядерный",
		"язык",
		"январь",
		"японский",
		"яркий",
		"ясный",
		"ящик",
	}
)
