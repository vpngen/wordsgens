package namesgenerator

var (
	rightNobelPeace = [...]Person{
		{
			Name:   "Анри Дюнан",
			Gender: 0,
			Desc:   "За вклад в мирное сотрудничество народов. Роль: основатель Международного комитета Красного Креста, один из создателей Женевской конвенции 1864 года.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%94%D1%8E%D0%BD%D0%B0%D0%BD,_%D0%90%D0%BD%D1%80%D0%B8",
		},
		{
			Name:   "Фредерик Пасси",
			Gender: 0,
			Desc:   "За многолетние миротворческие усилия. Роль: основатель и президент Лиги мира и свободы.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%B0%D1%81%D1%81%D0%B8,_%D0%A4%D1%80%D0%B5%D0%B4%D0%B5%D1%80%D0%B8%D0%BA",
		},
		{
			Name:   "Эли Дюкоммен",
			Gender: 0,
			Desc:   "В знак признания заслуг по основанию Международного бюро мира для координации деятельности различных пацифистских обществ Европы Роль: почётный секретарь Международного бюро мира (Берн, Швейцария)",
			URL:    "https://ru.wikipedia.org/wiki/%D0%94%D1%8E%D0%BA%D0%BE%D0%BC%D0%BC%D0%B5%D0%BD,_%D0%AD%D0%BB%D0%B8",
		},
		{
			Name:   "Шарль Гоба",
			Gender: 0,
			Desc:   "За усилия в деле международного арбитража Роль: почётный секретарь Международного бюро мира (Берн, Швейцария) и Межпарламентского союза",
			URL:    "https://ru.wikipedia.org/wiki/%D0%93%D0%BE%D0%B1%D0%B0,_%D0%A8%D0%B0%D1%80%D0%BB%D1%8C_%D0%90%D0%BB%D1%8C%D0%B1%D0%B5%D1%80",
		},
		{
			Name:   "Уильям Кример",
			Gender: 0,
			Desc:   "В ознаменование усилий по достижению мира путём арбитража. Роль: секретарь Международной лиги арбитража[en], член парламента Великобритании.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D1%80%D0%B8%D0%BC%D0%B5%D1%80,_%D0%A3%D0%B8%D0%BB%D1%8C%D1%8F%D0%BC_%D0%A0%D0%B0%D0%BD%D0%B4%D0%B0%D0%BB",
		},
		{
			Name:   "Берта Зуттнер",
			Gender: 1,
			Desc:   "За активную пацифистскую деятельность. Роль: автор «Lay Down Your Arms», почётный президент Международного бюро мира.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%97%D1%83%D1%82%D0%BD%D0%B5%D1%80,_%D0%91%D0%B5%D1%80%D1%82%D0%B0_%D1%84%D0%BE%D0%BD",
		},
		{
			Name:   "Теодор Рузвельт",
			Gender: 0,
			Desc:   "За свою роль в подписании Портсмутского договора. Роль: соавтор различных мирных договоров, президент Соединённых Штатов Америки.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D1%83%D0%B7%D0%B2%D0%B5%D0%BB%D1%8C%D1%82,_%D0%A2%D0%B5%D0%BE%D0%B4%D0%BE%D1%80",
		},
		{
			Name:   "Луи Рено",
			Gender: 0,
			Desc:   "Как подлинный гений международного права во Франции. Роль: профессор международного права.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D0%BD%D0%BE,_%D0%9B%D1%83%D0%B8_(%D1%8E%D1%80%D0%B8%D1%81%D1%82)",
		},
		{
			Name:   "Эрнесто Монета",
			Gender: 0,
			Desc:   "За свою неустанную деятельность во имя мира. Роль: президент Ломбардской лиги мира",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D0%BD%D0%B5%D1%82%D0%B0,_%D0%AD%D1%80%D0%BD%D0%B5%D1%81%D1%82%D0%BE_%D0%A2%D0%B5%D0%BE%D0%B4%D0%BE%D1%80%D0%BE",
		},
		{
			Name:   "Клас Арнольдсон",
			Gender: 0,
			Desc:   "За участие в разрешении норвежского конфликта. Роль: основатель Шведской лиги мира и арбитража[en], член парламента Швеции, писатель.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%80%D0%BD%D0%BE%D0%BB%D1%8C%D0%B4%D1%81%D0%BE%D0%BD,_%D0%9A%D0%BB%D0%B0%D1%81_%D0%9F%D0%BE%D0%BD%D1%82%D1%83%D1%81",
		},
		{
			Name:   "Фредерик Байер",
			Gender: 0,
			Desc:   "За создание Скандинавского межпарламентского союза для укрепления регионального сотрудничества. Роль: почётный президент Международного бюро мира, член датского парламента.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%B9%D0%B5%D1%80,_%D0%A4%D1%80%D0%B5%D0%B4%D1%80%D0%B8%D0%BA",
		},
		{
			Name:   "Огюст Беернарт",
			Gender: 0,
			Desc:   "В знак признания усилий в борьбе за международный арбитраж и сокращение вооружений. Роль: член Постоянной палаты третейского суда в Гааге, член парламента Бельгии, премьер-министр Бельгии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%B5%D1%80%D0%BD%D0%B0%D1%80%D1%82,_%D0%9E%D0%B3%D1%8E%D1%81%D1%82",
		},
		{
			Name:   "Поль Д’Эстурнель",
			Gender: 0,
			Desc:   "За договоры об арбитраже между Францией и соседними странами. Роль: основатель Комитета по защите национальных интересов и международного примирения, основатель и президент Французской парламентской группы добровольного арбитража, член парламента Франции.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AD%D1%82%D1%83%D1%80%D0%BD%D0%B5%D0%BB%D1%8C_%D0%B4%D0%B5_%D0%9A%D0%BE%D0%BD%D1%81%D1%82%D0%B0%D0%BD,_%D0%9F%D0%BE%D0%BB%D1%8C_%D0%90%D0%BD%D1%80%D0%B8_%D0%91%D0%B5%D0%BD%D0%B6%D0%B0%D0%BC%D0%B5%D0%BD_%D0%B4%E2%80%99",
		},
		{
			Name:   "Тобиас Ассер",
			Gender: 0,
			Desc:   "За работы в области международного арбитража. Роль: адвокат, член кабинета министров, инициатор Гаагской конференции по международному частному праву.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%81%D1%81%D0%B5%D1%80,_%D0%A2%D0%BE%D0%B1%D0%B8%D0%B0%D1%81",
		},
		{
			Name:   "Альфред Фрид",
			Gender: 0,
			Desc:   "За свою интернациональную деятельность. Роль: основатель Die Friedens-Warte, журналист.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A4%D1%80%D0%B8%D0%B4,_%D0%90%D0%BB%D1%8C%D1%84%D1%80%D0%B5%D0%B4_%D0%93%D0%B5%D1%80%D0%BC%D0%B0%D0%BD",
		},
		{
			Name:   "Элиу Рут",
			Gender: 0,
			Desc:   "За усилия по укреплению мира в западном полушарии. Роль: автор различных договоров арбитража, госсекретарь.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D1%83%D1%82,_%D0%AD%D0%BB%D0%B8%D1%83",
		},
		{
			Name:   "Анри Лафонтен",
			Gender: 0,
			Desc:   "Как истинный лидер народного движения за мир в Европе. Роль: президент Международного бюро мира, член парламента Бельгии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9B%D0%B0%D1%84%D0%BE%D0%BD%D1%82%D0%B5%D0%BD,_%D0%90%D0%BD%D1%80%D0%B8",
		},
		{
			Name:   "Вудро Вильсон",
			Gender: 0,
			Desc:   "За привнесение фундаментального закона человечности в современную международную политику. Роль: основатель Лиги Наций, президент Соединённых Штатов Америки",
			URL:    "https://ru.wikipedia.org/wiki/%D0%92%D0%B8%D0%BB%D1%8C%D1%81%D0%BE%D0%BD,_%D0%92%D1%83%D0%B4%D1%80%D0%BE",
		},
		{
			Name:   "Леон Буржуа",
			Gender: 0,
			Desc:   "За усилия по утверждению мира средствами арбитража. Роль: председатель Совета Лиги Наций, премьер-министр Франции.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D1%83%D1%80%D0%B6%D1%83%D0%B0,_%D0%9B%D0%B5%D0%BE%D0%BD",
		},
		{
			Name:   "Карл Брантинг",
			Gender: 0,
			Desc:   "За усилия в мирном решении спора Швеции с Норвегией и за работу в Лиге Наций. Роль: шведский делегат в Совете Лиги Наций, премьер-министр Швеции.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D1%80%D0%B0%D0%BD%D1%82%D0%B8%D0%BD%D0%B3,_%D0%9A%D0%B0%D1%80%D0%BB_%D0%AF%D0%BB%D1%8C%D0%BC%D0%B0%D1%80",
		},
		{
			Name:   "Кристиан Ланге",
			Gender: 0,
			Desc:   "За пропаганду арбитража как средства решения международных конфликтов. Роль: генеральный секретарь Межпарламентского союза.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9B%D0%B0%D0%BD%D0%B3%D0%B5,_%D0%9A%D1%80%D0%B8%D1%81%D1%82%D0%B8%D0%B0%D0%BD_%D0%9B%D0%BE%D1%83%D1%81",
		},
		{
			Name:   "Фритьоф Нансен",
			Gender: 0,
			Desc:   "За многолетние усилия по оказанию помощи беззащитным. Роль: создатель Нансеновского паспорта для беженцев, норвежский делегат в Лиге Наций, исследователь, учёный.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9D%D0%B0%D0%BD%D1%81%D0%B5%D0%BD,_%D0%A4%D1%80%D0%B8%D1%82%D1%8C%D0%BE%D1%84",
		},
		{
			Name:   "Джозеф Чемберлен",
			Gender: 0,
			Desc:   "За свою роль в локарнских переговорах. Роль: участник локарнских переговоров, министр иностранных дел Великобритании.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A7%D0%B5%D0%BC%D0%B1%D0%B5%D1%80%D0%BB%D0%B5%D0%BD,_%D0%94%D0%B6%D0%BE%D0%B7%D0%B5%D1%84_%D0%9E%D1%81%D1%82%D0%B8%D0%BD",
		},
		{
			Name:   "Чарлз Дауэс",
			Gender: 0,
			Desc:   "В знак признания вклада в план, носящий его имя. Роль: председатель Комиссии по репарациям, вице-президент Соединённых Штатов Америки.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%94%D0%B0%D1%83%D1%8D%D1%81,_%D0%A7%D0%B0%D1%80%D0%BB%D0%B7_%D0%93%D0%B5%D0%B9%D1%82%D1%81",
		},
		{
			Name:   "Аристид Бриан",
			Gender: 0,
			Desc:   "За роль в заключении Локарнского пакта и дружественном диалоге Франции и Германии после многих лет недоверия. Роль: соавтор Локарнского пакта и пакта Бриана — Келлога, министр иностранных дел Франции.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D1%80%D0%B8%D0%B0%D0%BD,_%D0%90%D1%80%D0%B8%D1%81%D1%82%D0%B8%D0%B4",
		},
		{
			Name:   "Густав Штреземан",
			Gender: 0,
			Desc:   "За роль в заключении Локарнского пакта и дружественном диалоге Франции и Германии после многих лет недоверия. Роль: соавтор Локарнского пакта, министр иностранных дел Германии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A8%D1%82%D1%80%D0%B5%D0%B7%D0%B5%D0%BC%D0%B0%D0%BD,_%D0%93%D1%83%D1%81%D1%82%D0%B0%D0%B2",
		},
		{
			Name:   "Фердинанд Бюиссон",
			Gender: 0,
			Desc:   "За деятельность, направленную на восстановление понимания между французским и германским народами. Роль: основатель и президент Лиги по правам человека, профессор Сорбонны.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D1%8E%D0%B8%D1%81%D1%81%D0%BE%D0%BD,_%D0%A4%D0%B5%D1%80%D0%B4%D0%B8%D0%BD%D0%B0%D0%BD%D0%B4_%D0%AD%D0%B4%D1%83%D0%B0%D1%80%D0%B4",
		},
		{
			Name:   "Людвиг Квидде",
			Gender: 0,
			Desc:   "За подготовку общественного мнения Франции и Германии к мирному сотрудничеству. Роль: участник различных мирных конференций, член парламента Германии, профессор Берлинского университета.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B2%D0%B8%D0%B4%D0%B4%D0%B5,_%D0%9B%D1%8E%D0%B4%D0%B2%D0%B8%D0%B3",
		},
		{
			Name:   "Фрэнк Келлог",
			Gender: 0,
			Desc:   "За подготовку Парижского пакта. Роль: соавтор Парижского пакта, государственный секретарь США.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B5%D0%BB%D0%BB%D0%BE%D0%B3,_%D0%A4%D1%80%D1%8D%D0%BD%D0%BA_%D0%91%D0%B8%D0%BB%D0%BB%D0%B8%D0%BD%D0%B3%D1%81",
		},
		{
			Name:   "Натан Сёдерблум",
			Gender: 0,
			Desc:   "В ознаменование заслуг в достижении мира через религиозное объединение. Роль: лидер экуменического движения, архиепископ.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D1%91%D0%B4%D0%B5%D1%80%D0%B1%D0%BB%D1%83%D0%BC,_%D0%9D%D0%B0%D1%82%D0%B0%D0%BD",
		},
		{
			Name:   "Джейн Аддамс",
			Gender: 1,
			Desc:   "Как подлинный делегат всех миролюбивых женщин мира. Роль: президент Международной женской лиги за мир и свободу, социолог",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D0%B4%D0%B4%D0%B0%D0%BC%D1%81,_%D0%94%D0%B6%D0%B5%D0%B9%D0%BD",
		},
		{
			Name:   "Николас Батлер",
			Gender: 0,
			Desc:   "За неиссякаемую энергию и рвение в деле мира. Роль: сторонник пакта Бриана — Келлога, президент Колумбийского университета.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D1%82%D0%BB%D0%B5%D1%80,_%D0%9D%D0%B8%D0%BA%D0%BE%D0%BB%D0%B0%D1%81_%D0%9C%D1%8E%D1%80%D1%80%D0%B5%D0%B9",
		},
		{
			Name:   "Норман Эйнджелл",
			Gender: 0,
			Desc:   "За пропаганду мира. Роль: член комиссии Исполнительного комитета Лиги Наций и Национального совета мира, писатель.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AD%D0%B9%D0%BD%D0%B4%D0%B6%D0%B5%D0%BB%D0%BB,_%D0%9D%D0%BE%D1%80%D0%BC%D0%B0%D0%BD",
		},
		{
			Name:   "Артур Хендерсон",
			Gender: 0,
			Desc:   "За настойчивую защиту дела международного разоружения. Роль: Председатель Конференции по разоружению в 1932 году, министр иностранных дел Великобритании.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A5%D0%B5%D0%BD%D0%B4%D0%B5%D1%80%D1%81%D0%BE%D0%BD,_%D0%90%D1%80%D1%82%D1%83%D1%80",
		},
		{
			Name:   "Карл Осецкий",
			Gender: 0,
			Desc:   "За борьбу с милитаризмом в Германии. Роль: пацифист, журналист.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9E%D1%81%D0%B5%D1%86%D0%BA%D0%B8%D0%B9,_%D0%9A%D0%B0%D1%80%D0%BB_%D1%84%D0%BE%D0%BD",
		},
		{
			Name:   "Карлос Ламас",
			Gender: 0,
			Desc:   "За миротворческую роль в боливийско-парагвайском конфликте 1935 года. Роль: министр иностранных дел Аргентины, председатель Ассамблеи Лиги Наций, посредник в конфликте между Парагваем и Боливией.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D0%B0%D0%B2%D0%B5%D0%B4%D1%80%D0%B0_%D0%9B%D0%B0%D0%BC%D0%B0%D1%81,_%D0%9A%D0%B0%D1%80%D0%BB%D0%BE%D1%81",
		},
		{
			Name:   "Роберт Сесил",
			Gender: 0,
			Desc:   "В ознаменование заслуг перед Лигой Наций. Роль: основатель и президент Международной кампании мира, писатель, лорд-хранитель Малой печати.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B5%D1%81%D0%B8%D0%BB,_%D0%A0%D0%BE%D0%B1%D0%B5%D1%80%D1%82_(%D1%8E%D1%80%D0%B8%D1%81%D1%82)",
		},
		{
			Name:   "Корделл Халл",
			Gender: 0,
			Desc:   "В знак признания его заслуг по утверждению мира в западном полушарии, в укреплении торговли и становлении ООН. Роль: важный участник создания Организации Объединённых Наций, госсекретарь США.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A5%D0%B0%D0%BB%D0%BB,_%D0%9A%D0%BE%D1%80%D0%B4%D0%B5%D0%BB%D0%BB",
		},
		{
			Name:   "Эмили Болч",
			Gender: 1,
			Desc:   "За многолетний, неутомимый труд на благо мира. Роль: почётный президент Международной лиги женщин за мир и свободу, профессор истории и социологии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D0%BB%D1%87,_%D0%AD%D0%BC%D0%B8%D0%BB%D0%B8_%D0%93%D1%80%D0%B8%D0%BD",
		},
		{
			Name:   "Джон Мотт",
			Gender: 0,
			Desc:   "За миссионерскую деятельность. Роль: президент Всемирного альянса ассоциаций молодых христиан, председатель Международного миссионерского совета.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D1%82%D1%82,_%D0%94%D0%B6%D0%BE%D0%BD_%D0%A0%D1%8D%D0%BB%D0%B5%D0%B9",
		},
		{
			Name:   "Джон Бойд Орр",
			Gender: 0,
			Desc:   "В знак признания его заслуг не только в деле освобождения человечества от нужды, но и в создании основ мирной кооперации между классами, нациями и расами. Роль: президент Национального совета по вопросам мира и Всемирного союза организаций мира, основатель и директор Всеобщей продовольственной и сельскохозяйственной организации, политик, врач",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D0%B9%D0%B4_%D0%9E%D1%80%D1%80,_%D0%94%D0%B6%D0%BE%D0%BD",
		},
		{
			Name:   "Ральф Банч",
			Gender: 0,
			Desc:   "За создание Скандинавского межпарламентского союза для укрепления регионального сотрудничества. Роль: посредник в Палестине, директор отдела по опеке ООН, профессор Гарвардского университета.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B0%D0%BD%D1%87,_%D0%A0%D0%B0%D0%BB%D1%8C%D1%84",
		},
		{
			Name:   "Леон Жуо",
			Gender: 0,
			Desc:   "В ознаменование миротворческих заслуг. Роль: делегат ООН, член Совета Международной организации труда, вице-президент Всемирной федерации профсоюзов, вице-президент Международной конфедерации свободных профсоюзов, президент Национального экономического совета и Международного комитета европейского совета, президент конфедерации профсоюзов «C.G.T. Force Ouvrière».",
			URL:    "https://ru.wikipedia.org/wiki/%D0%96%D1%83%D0%BE,_%D0%9B%D0%B5%D0%BE%D0%BD",
		},
		{
			Name:   "Альберт Швейцер",
			Gender: 0,
			Desc:   "За миссионерскую деятельность. Роль: основатель больницы в Ламбарене, миссионерский хирург.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A8%D0%B2%D0%B5%D0%B9%D1%86%D0%B5%D1%80,_%D0%90%D0%BB%D1%8C%D0%B1%D0%B5%D1%80%D1%82",
		},
		{
			Name:   "Джордж Маршалл",
			Gender: 0,
			Desc:   "Как инициатор «Плана Маршалла» по послевоенному восстановлению экономики Европы. Роль: инициатор «плана Маршалла», делегат ООН, госсекретарь и министр обороны США, президент Американского Красного Креста.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D1%80%D1%88%D0%B0%D0%BB%D0%BB,_%D0%94%D0%B6%D0%BE%D1%80%D0%B4%D0%B6",
		},
		{
			Name:   "Лестер Пирсон",
			Gender: 0,
			Desc:   "За свою роль в преодолении Суэцкого кризиса. Роль: президент 7-й сессии Генеральной Ассамблеи ООН, государственный секретарь по вопросам иностранных дел Канады",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%B8%D1%80%D1%81%D0%BE%D0%BD,_%D0%9B%D0%B5%D1%81%D1%82%D0%B5%D1%80",
		},
		{
			Name:   "Жорж Пир",
			Gender: 0,
			Desc:   "За помощь беженцам. Роль: руководитель организации по оказанию помощи беженцам «l’Europe du Coeur au Service du Monde», глава ордена доминиканцев.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%B8%D1%80,_%D0%96%D0%BE%D1%80%D0%B6",
		},
		{
			Name:   "Филип Ноэль-Бейкер",
			Gender: 0,
			Desc:   "Как крупнейший специалист по разоружению. Роль: сторонник международного мира и сотрудничества, член парламента Великобритании.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D1%8D%D0%BB%D1%8C-%D0%91%D0%B5%D0%B9%D0%BA%D0%B5%D1%80,_%D0%A4%D0%B8%D0%BB%D0%B8%D0%BF",
		},
		{
			Name:   "Альберт Лутули",
			Gender: 0,
			Desc:   "За усилия по утверждению справедливости между людьми и народами. Роль: президент Африканского национального конгресса в ЮАР.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9B%D1%83%D1%82%D1%83%D0%BB%D0%B8,_%D0%90%D0%BB%D1%8C%D0%B1%D0%B5%D1%80%D1%82",
		},
		{
			Name:   "Даг Хаммаршёльд",
			Gender: 0,
			Desc:   "За деятельность в ООН. Роль: генеральный секретарь ООН.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A5%D0%B0%D0%BC%D0%BC%D0%B0%D1%80%D1%88%D1%91%D0%BB%D1%8C%D0%B4,_%D0%94%D0%B0%D0%B3",
		},
		{
			Name:   "Лайнус Полинг",
			Gender: 0,
			Desc:   "Как автор проекта договора о запрещении ядерных испытаний.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%BE%D0%BB%D0%B8%D0%BD%D0%B3,_%D0%9B%D0%B0%D0%B9%D0%BD%D1%83%D1%81",
		},
		{
			Name:   "Мартин Лютер Кинг",
			Gender: 0,
			Desc:   "За деятельность в пользу равноправия темнокожих. Роль: лидер «Южной конференции христианских лидеров».",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D0%BD%D0%B3,_%D0%9C%D0%B0%D1%80%D1%82%D0%B8%D0%BD_%D0%9B%D1%8E%D1%82%D0%B5%D1%80",
		},
		{
			Name:   "Рене Кассен",
			Gender: 0,
			Desc:   "В ознаменование 20-й годовщины принятия Декларации прав человека. Роль: председатель Европейского суда по правам человека.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D1%81%D1%81%D0%B5%D0%BD,_%D0%A0%D0%B5%D0%BD%D0%B5",
		},
		{
			Name:   "Норман Борлоуг",
			Gender: 0,
			Desc:   "За вклад в решение продовольственной проблемы, и особенно за осуществление «зелёной революции». Роль: участник Международного центра улучшения кукурузы и пшеницы.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%BE%D1%80%D0%BB%D0%BE%D1%83%D0%B3,_%D0%9D%D0%BE%D1%80%D0%BC%D0%B0%D0%BD_%D0%AD%D1%80%D0%BD%D0%B5%D1%81%D1%82",
		},
		{
			Name:   "Вилли Брандт",
			Gender: 0,
			Desc:   "В знак признания конкретных инициатив, повлёкших ослабление напряжённости между Востоком и Западом. Роль: канцлер ФРГ.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D1%80%D0%B0%D0%BD%D0%B4%D1%82,_%D0%92%D0%B8%D0%BB%D0%BB%D0%B8",
		},
		{
			Name:   "Генри Киссинджер",
			Gender: 0,
			Desc:   "В знак признания заслуг в связи с перемирием во Вьетнаме. Роль: государственный секретарь США.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D1%81%D1%81%D0%B8%D0%BD%D0%B4%D0%B6%D0%B5%D1%80,_%D0%93%D0%B5%D0%BD%D1%80%D0%B8",
		},
		{
			Name:   "Ле Дык Тхо",
			Gender: 0,
			Desc:   "В знак признания заслуг в связи с перемирием во Вьетнаме.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9B%D0%B5_%D0%94%D1%8B%D0%BA_%D0%A2%D1%85%D0%BE",
		},
		{
			Name:   "Шон Макбрайд",
			Gender: 0,
			Desc:   "За создание международных механизмов наблюдения за состоянием прав человека. Роль: председатель комиссии по Намибии ООН, президент Международного бюро мира.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%BA%D0%B1%D1%80%D0%B0%D0%B9%D0%B4,_%D0%A8%D0%BE%D0%BD",
		},
		{
			Name:   "Эйсаку Сато",
			Gender: 0,
			Desc:   "За политику антимилитаризма. Роль: премьер-министр Японии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D1%82%D0%BE,_%D0%AD%D0%B9%D1%81%D0%B0%D0%BA%D1%83",
		},
		{
			Name:   "Андрей Сахаров",
			Gender: 0,
			Desc:   "За бесстрашную поддержку фундаментальных принципов мира между людьми и мужественную борьбу со злоупотреблением властью и любыми формами подавления человеческого достоинства. Роль: советский физик-ядерщик.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D1%85%D0%B0%D1%80%D0%BE%D0%B2,_%D0%90%D0%BD%D0%B4%D1%80%D0%B5%D0%B9_%D0%94%D0%BC%D0%B8%D1%82%D1%80%D0%B8%D0%B5%D0%B2%D0%B8%D1%87",
		},
		{
			Name:   "Бетти Уильямс",
			Gender: 1,
			Desc:   "В знак признания заслуг в деле мира. Роль: основательницы движения за мир в Северной Ирландии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A3%D0%B8%D0%BB%D1%8C%D1%8F%D0%BC%D1%81,_%D0%91%D0%B5%D1%82%D1%82%D0%B8",
		},
		{
			Name:   "Мейрид Корриган",
			Gender: 1,
			Desc:   "В знак признания заслуг в деле мира. Роль: основательницы движения за мир в Северной Ирландии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%BE%D1%80%D1%80%D0%B8%D0%B3%D0%B0%D0%BD,_%D0%9C%D0%B5%D0%B9%D1%80%D0%B8%D0%B4",
		},
		{
			Name:   "Анвар Садат",
			Gender: 0,
			Desc:   "За подготовку и заключение основополагающих соглашений между Израилем и Египтом. Роль: президент Египта.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D0%B4%D0%B0%D1%82,_%D0%90%D0%BD%D0%B2%D0%B0%D1%80",
		},
		{
			Name:   "Менахем Бегин",
			Gender: 0,
			Desc:   "За подготовку и заключение основополагающих соглашений между Израилем и Египтом. Роль: премьер-министр Израиля.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%B3%D0%B8%D0%BD,_%D0%9C%D0%B5%D0%BD%D0%B0%D1%85%D0%B5%D0%BC",
		},
		{
			Name:   "Мать Тереза",
			Gender: 1,
			Desc:   "Утверждает мир в самой важной сфере, защищая неприкосновенность человеческого достоинства. Роль: основатель организации «Сёстры — миссионерки любви».",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D1%82%D1%8C_%D0%A2%D0%B5%D1%80%D0%B5%D0%B7%D0%B0",
		},
		{
			Name:   "Адольфо Эскивель",
			Gender: 0,
			Desc:   "Как борец за права человека. Роль: лидер в области прав человека, скульптор, архитектор.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%B5%D1%80%D0%B5%D1%81_%D0%AD%D1%81%D0%BA%D0%B8%D0%B2%D0%B5%D0%BB%D1%8C,_%D0%90%D0%B4%D0%BE%D0%BB%D1%8C%D1%84%D0%BE",
		},
		{
			Name:   "Альва Мюрдаль",
			Gender: 1,
			Desc:   "За крупный вклад в дело разоружения. Роль: писатель, дипломат, член кабинета министров.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D1%8E%D1%80%D0%B4%D0%B0%D0%BB%D1%8C,_%D0%90%D0%BB%D1%8C%D0%B2%D0%B0",
		},
		{
			Name:   "Альфонсо Роблес",
			Gender: 0,
			Desc:   "За крупный вклад в дело разоружения. Роль: министр иностранных дел, делегат Генеральной Ассамблеи Организации Объединённых Наций по разоружению, дипломат.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%93%D0%B0%D1%80%D1%81%D0%B8%D1%8F_%D0%A0%D0%BE%D0%B1%D0%BB%D0%B5%D1%81,_%D0%90%D0%BB%D1%8C%D1%84%D0%BE%D0%BD%D1%81%D0%BE",
		},
		{
			Name:   "Лех Валенса",
			Gender: 0,
			Desc:   "Как борец за права человека. Роль: профсоюзный лидер.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%92%D0%B0%D0%BB%D0%B5%D0%BD%D1%81%D0%B0,_%D0%9B%D0%B5%D1%85",
		},
		{
			Name:   "Десмонд Туту",
			Gender: 0,
			Desc:   "За мужество и героизм, проявленные в борьбе против апартеида. Роль: генеральный секретарь Совета церквей ЮАР, епископ Йоханнесбурга.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A2%D1%83%D1%82%D1%83,_%D0%94%D0%B5%D1%81%D0%BC%D0%BE%D0%BD%D0%B4",
		},
		{
			Name:   "Эли Визель",
			Gender: 0,
			Desc:   "За приверженность тематике, посвящённой страданиям еврейского народа, жертвам нацизма. Роль: председатель «Президентской комиссии по Холокосту».",
			URL:    "https://ru.wikipedia.org/wiki/%D0%92%D0%B8%D0%B7%D0%B5%D0%BB%D1%8C,_%D0%AD%D0%BB%D0%B8",
		},
		{
			Name:   "Оскар Санчес",
			Gender: 0,
			Desc:   "ак инициатор мирных переговоров в Центральной Америке. Роль: президент Коста-Рики.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%80%D0%B8%D0%B0%D1%81_%D0%A1%D0%B0%D0%BD%D1%87%D0%B5%D1%81,_%D0%9E%D1%81%D0%BA%D0%B0%D1%80",
		},
		{
			Name:   "14-й Далай-лама",
			Gender: 0,
			Desc:   "За неустанную проповедь добросердечия, любви и терпимости в отношениях между отдельными людьми, сообществами и народами.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%94%D0%B0%D0%BB%D0%B0%D0%B9-%D0%BB%D0%B0%D0%BC%D0%B0_XIV",
		},
		{
			Name:   "Михаил Горбачёв",
			Gender: 0,
			Desc:   "В знак признания его ведущей роли в мирном процессе, который сегодня характеризует важную составную часть жизни международного сообщества. Роль: президент СССР.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%93%D0%BE%D1%80%D0%B1%D0%B0%D1%87%D1%91%D0%B2,_%D0%9C%D0%B8%D1%85%D0%B0%D0%B8%D0%BB_%D0%A1%D0%B5%D1%80%D0%B3%D0%B5%D0%B5%D0%B2%D0%B8%D1%87",
		},
		{
			Name:   "Аун Сан Су Чжи",
			Gender: 1,
			Desc:   "За её ненасильственную борьбу за демократию и права человека.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%83%D0%BD_%D0%A1%D0%B0%D0%BD_%D0%A1%D1%83_%D0%A7%D0%B6%D0%B8",
		},
		{
			Name:   "Ригоберта Менчу",
			Gender: 1,
			Desc:   "В знак признания её работы за социальную справедливость и этнокультурное примирение на основе уважения прав коренных народов.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D0%BD%D1%87%D1%83,_%D0%A0%D0%B8%D0%B3%D0%BE%D0%B1%D0%B5%D1%80%D1%82%D0%B0",
		},
		{
			Name:   "Нельсон Мандела",
			Gender: 0,
			Desc:   "За работу по мирному прекращению режима апартеида, и за подготовку основы для новой демократической Южно-Африканской Республики.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%BD%D0%B4%D0%B5%D0%BB%D0%B0,_%D0%9D%D0%B5%D0%BB%D1%8C%D1%81%D0%BE%D0%BD",
		},
		{
			Name:   "Фредерик де Клерк",
			Gender: 0,
			Desc:   "За работу по мирному прекращению режима апартеида, и за подготовку основы для новой демократической Южно-Африканской Республики.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%BB%D0%B5%D1%80%D0%BA,_%D0%A4%D1%80%D0%B5%D0%B4%D0%B5%D1%80%D0%B8%D0%BA_%D0%92%D0%B8%D0%BB%D0%BB%D0%B5%D0%BC_%D0%B4%D0%B5",
		},
		{
			Name:   "Ясир Арафат",
			Gender: 0,
			Desc:   "За усилия по достижению мира на Ближнем Востоке. Роль: председатель Палестинской национальной администрации, председатель исполнительного комитета Организации освобождения Палестины.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%80%D0%B0%D1%84%D0%B0%D1%82,_%D0%AF%D1%81%D0%B8%D1%80",
		},
		{
			Name:   "Шимон Перес",
			Gender: 0,
			Desc:   "За усилия по достижению мира на Ближнем Востоке. Роль: министр иностранных дел Израиля.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9F%D0%B5%D1%80%D0%B5%D1%81,_%D0%A8%D0%B8%D0%BC%D0%BE%D0%BD",
		},
		{
			Name:   "Ицхак Рабин",
			Gender: 0,
			Desc:   "За усилия по достижению мира на Ближнем Востоке. Роль: премьер-министр Израиля",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%B1%D0%B8%D0%BD,_%D0%98%D1%86%D1%85%D0%B0%D0%BA",
		},
		{
			Name:   "Джозеф Ротблат",
			Gender: 0,
			Desc:   "За усилия, направленные на снижение роли ядерного оружия в международной политике и, в долгосрочной перспективе, за запрещение этого вида оружия.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D0%BE%D1%82%D0%B1%D0%BB%D0%B0%D1%82,_%D0%94%D0%B6%D0%BE%D0%B7%D0%B5%D1%84",
		},
		{
			Name:   "Карлуш Белу",
			Gender: 0,
			Desc:   "За усилия по справедливому и мирному разрешению конфликта в Восточном Тиморе.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%BB%D1%83,_%D0%9A%D0%B0%D1%80%D0%BB%D1%83%D1%88",
		},
		{
			Name:   "Жозе Рамуш-Орта",
			Gender: 0,
			Desc:   "За усилия по справедливому и мирному разрешению конфликта в Восточном Тиморе.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%BC%D1%83%D1%88-%D0%9E%D1%80%D1%82%D0%B0,_%D0%96%D0%BE%D0%B7%D0%B5",
		},
		{
			Name:   "Джоди Уильямс",
			Gender: 1,
			Desc:   "За работу по запрещению и ликвидации противопехотных мин.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A3%D0%B8%D0%BB%D1%8C%D1%8F%D0%BC%D1%81,_%D0%94%D0%B6%D0%BE%D0%B4%D0%B8",
		},
		{
			Name:   "Джон Хьюм",
			Gender: 0,
			Desc:   "За усилия по поиску мирного решения конфликта в Северной Ирландии. Роль: депутат, лидер Социал-демократической и лейбористской партии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A5%D1%8C%D1%8E%D0%BC,_%D0%94%D0%B6%D0%BE%D0%BD",
		},
		{
			Name:   "Дэвид Тримбл",
			Gender: 0,
			Desc:   "За усилия по поиску мирного решения конфликта в Северной Ирландии. Роль: депутат, лидер Ольстерской юнионистской партии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A2%D1%80%D0%B8%D0%BC%D0%B1%D0%BB,_%D0%94%D1%8D%D0%B2%D0%B8%D0%B4",
		},
		{
			Name:   "Ким Дэ Чжун",
			Gender: 0,
			Desc:   "За работу над проблемами демократии и прав человека в Южной Корее и в Восточной Азии в целом, а также за работу по примирению с Северной Кореей в частности. Роль: президент Южной Кореи.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B8%D0%BC_%D0%94%D1%8D_%D0%A7%D0%B6%D1%83%D0%BD",
		},
		{
			Name:   "Кофи Аннан",
			Gender: 0,
			Desc:   "За вклад в создание более организованного мира и укрепление мира во всем мире. Роль: генеральный секретарь ООН.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D0%BD%D0%BD%D0%B0%D0%BD,_%D0%9A%D0%BE%D1%84%D0%B8",
		},
		{
			Name:   "Джимми Картер",
			Gender: 0,
			Desc:   "За десятилетия работы по поиску мирных решений международных конфликтов, укрепление демократии и прав человека, а также содействие экономическому и социальному развитию. Роль: 39-й президент Соединённых Штатов Америки.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D1%80%D1%82%D0%B5%D1%80,_%D0%94%D0%B6%D0%B8%D0%BC%D0%BC%D0%B8",
		},
		{
			Name:   "Ширин Эбади",
			Gender: 1,
			Desc:   "За вклад в развитие демократии и борьбу за права человека, особенно женщин и детей в Иране.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AD%D0%B1%D0%B0%D0%B4%D0%B8,_%D0%A8%D0%B8%D1%80%D0%B8%D0%BD",
		},
		{
			Name:   "Вангари Маатаи",
			Gender: 1,
			Desc:   "За вклад в устойчивое развитие, демократию и мир.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%B0%D0%B0%D1%82%D0%B0%D0%B8,_%D0%92%D0%B0%D0%BD%D0%B3%D0%B0%D1%80%D0%B8",
		},
		{
			Name:   "Мохаммед аль-Барадеи",
			Gender: 0,
			Desc:   "За усилия по предотвращению использования атомной энергии в военных целях и по обеспечению её применения в мирных целях в максимально безопасных условиях.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AD%D0%BB%D1%8C-%D0%91%D0%B0%D1%80%D0%B0%D0%B4%D0%B5%D0%B8,_%D0%9C%D0%BE%D1%85%D0%B0%D0%BC%D0%BC%D0%B5%D0%B4",
		},
		{
			Name:   "Мухаммад Юнус",
			Gender: 0,
			Desc:   "За усилия по созданию основ для социального и экономического развития. Роль: основатель Grameen Bank.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AE%D0%BD%D1%83%D1%81,_%D0%9C%D1%83%D1%85%D0%B0%D0%BC%D0%BC%D0%B0%D0%B4",
		},
		{
			Name:   "Альберт Гор",
			Gender: 0,
			Desc:   "За изучение последствий глобальных климатических изменений, вызванных деятельностью человека, и выработку мер по их возможному предотвращению.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%93%D0%BE%D1%80,_%D0%90%D0%BB%D1%8C%D0%B1%D0%B5%D1%80%D1%82",
		},
		{
			Name:   "Мартти Ахтисаари",
			Gender: 0,
			Desc:   "За те важные усилия в разрешении международных конфликтов, которые он прилагал на нескольких континентах в течение трёх десятилетий.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D1%85%D1%82%D0%B8%D1%81%D0%B0%D0%B0%D1%80%D0%B8,_%D0%9C%D0%B0%D1%80%D1%82%D1%82%D0%B8",
		},
		{
			Name:   "Барак Обама",
			Gender: 0,
			Desc:   "За выдающиеся усилия по укреплению международной дипломатии и сотрудничества между народами. Роль: 44-й президент Соединённых Штатов Америки.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D0%B0%D0%BC%D0%B0,_%D0%91%D0%B0%D1%80%D0%B0%D0%BA",
		},
		{
			Name:   "Лю Сяобо",
			Gender: 0,
			Desc:   "За длительную ненасильственную борьбу за фундаментальные права человека в Китае.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BB%D0%B0%D1%83%D1%80%D0%B5%D0%B0%D1%82%D0%BE%D0%B2_%D0%9D%D0%BE%D0%B1%D0%B5%D0%BB%D0%B5%D0%B2%D1%81%D0%BA%D0%BE%D0%B9_%D0%BF%D1%80%D0%B5%D0%BC%D0%B8%D0%B8_%D0%BC%D0%B8%D1%80%D0%B0#:~:text=2010-,%D0%9B%D1%8E%20%D0%A1%D1%8F%D0%BE%D0%B1%D0%BE,-(1955%E2%80%942017)",
		},
		{
			Name:   "Элен Джонсон-Серлиф",
			Gender: 1,
			Desc:   "За ненасильственную борьбу за права и безопасность женщин и участие в миротворческом процессе.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%94%D0%B6%D0%BE%D0%BD%D1%81%D0%BE%D0%BD-%D0%A1%D0%B5%D1%80%D0%BB%D0%B8%D1%84,_%D0%AD%D0%BB%D0%B5%D0%BD",
		},
		{
			Name:   "Лейма Гбови",
			Gender: 1,
			Desc:   "За ненасильственную борьбу за права и безопасность женщин и участие в миротворческом процессе.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%93%D0%B1%D0%BE%D0%B2%D0%B8,_%D0%9B%D0%B5%D0%B9%D0%BC%D0%B0",
		},
		{
			Name:   "Тавакуль Карман",
			Gender: 1,
			Desc:   "За ненасильственную борьбу за права и безопасность женщин и участие в миротворческом процессе.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9A%D0%B0%D1%80%D0%BC%D0%B0%D0%BD,_%D0%A2%D0%B0%D0%B2%D0%B0%D0%BA%D1%83%D0%BB%D1%8C",
		},
		{
			Name:   "Малала Юсуфзай",
			Gender: 1,
			Desc:   "За борьбу против притеснений детей и молодёжи и за права детей на образование.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%AE%D1%81%D1%83%D1%84%D0%B7%D0%B0%D0%B9,_%D0%9C%D0%B0%D0%BB%D0%B0%D0%BB%D0%B0",
		},
		{
			Name:   "Кайлаш Сатьяртхи",
			Gender: 0,
			Desc:   "За борьбу против притеснений детей и молодёжи и за права детей на образование.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D1%82%D1%8C%D1%8F%D1%80%D1%82%D1%85%D0%B8,_%D0%9A%D0%B0%D0%B9%D0%BB%D0%B0%D1%88",
		},
		{
			Name:   "Хуан Сантос",
			Gender: 0,
			Desc:   "За усилия по прекращению в стране более чем полувековой гражданской войны. Роль: президент Колумбии.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A1%D0%B0%D0%BD%D1%82%D0%BE%D1%81,_%D0%A5%D1%83%D0%B0%D0%BD_%D0%9C%D0%B0%D0%BD%D1%83%D1%8D%D0%BB%D1%8C",
		},
		{
			Name:   "Дени Муквеге",
			Gender: 0,
			Desc:   "За попытку прекратить использование сексуального насилия в качестве оружия в войнах и вооружённых конфликтах.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D1%83%D0%BA%D0%B2%D0%B5%D0%B3%D0%B5,_%D0%94%D0%B5%D0%BD%D0%B8",
		},
		{
			Name:   "Надя Мурад",
			Gender: 1,
			Desc:   "За попытку прекратить использование сексуального насилия в качестве оружия в войнах и вооружённых конфликтах.",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D1%83%D1%80%D0%B0%D0%B4,_%D0%9D%D0%B0%D0%B4%D1%8F",
		},
		{
			Name:   "Абий Ахмед Али",
			Gender: 0,
			Desc:   "За разрешение пограничного конфликта с Эритреей",
			URL:    "https://ru.wikipedia.org/wiki/%D0%90%D0%B1%D0%B8%D0%B9_%D0%90%D1%85%D0%BC%D0%B5%D0%B4_%D0%90%D0%BB%D0%B8",
		},
		{
			Name:   "Мария Ресса",
			Gender: 1,
			Desc:   "За усилия по защите свободы слова, которая является необходимым условием демократии и прочного мира",
			URL:    "https://ru.wikipedia.org/wiki/%D0%A0%D0%B5%D1%81%D1%81%D0%B0,_%D0%9C%D0%B0%D1%80%D0%B8%D1%8F",
		},
		{
			Name:   "Дмитрий Муратов",
			Gender: 0,
			Desc:   "За усилия по защите свободы слова, которая является необходимым условием демократии и прочного мира",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D1%83%D1%80%D0%B0%D1%82%D0%BE%D0%B2,_%D0%94%D0%BC%D0%B8%D1%82%D1%80%D0%B8%D0%B9_%D0%90%D0%BD%D0%B4%D1%80%D0%B5%D0%B5%D0%B2%D0%B8%D1%87",
		},
		{
			Name:   "Алесь Беляцкий",
			Gender: 0,
			Desc:   "Лауреаты Премии мира представляют гражданское общество в своих странах. Они на протяжении многих лет продвигали право критиковать власть и защищать основные права граждан. Они приложили выдающиеся усилия, чтобы задокументировать военные преступления, нарушения прав человека и злоупотребления властью. Вместе они демонстрируют важность гражданского общества для мира и демократии",
			URL:    "https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D0%BB%D1%8F%D1%86%D0%BA%D0%B8%D0%B9,_%D0%90%D0%BB%D0%B5%D1%81%D1%8C_%D0%92%D0%B8%D0%BA%D1%82%D0%BE%D1%80%D0%BE%D0%B2%D0%B8%D1%87",
		},
		{
			Name:   "Наргиз Сафие Мохаммади",
			Gender: 1,
			Desc:   "За борьбу против угнетения женщин в Иране и за продвижение прав человека и свободы для всех",
			URL:    "https://ru.wikipedia.org/wiki/%D0%9C%D0%BE%D1%85%D0%B0%D0%BC%D0%BC%D0%B0%D0%B4%D0%B8,_%D0%9D%D0%B0%D1%80%D0%B3%D0%B8%D0%B7",
		},
	}
)
