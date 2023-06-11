const cities = [
    {
        img: './src/assets/img/cities/yalta.png', 
        name: 'Ялта',
        count: '134 отеля',
    },
    {
        img: './src/assets/img/cities/sudak.png', 
        name: 'Судак',
        count: '113 отелей',
    },
    {
        img: './src/assets/img/cities/balaklava.png', 
        name: 'Балаклава',
        count: '13 отелей',
    },
    {
        img: './src/assets/img/cities/alushta.png', 
        name: 'Алушта',
        count: '111 отелей',
    }
]

export function PopDir(){
    return(
        <>
            <div className="section_title">
                <div className="name_title">Популярные направления</div>
                <div className="more">
                    <div className="more_title">Смотреть больше</div>
                    <div className="more_img">
                        <img src="./src/assets/img/components/arrow_more.png" alt="" />
                    </div>
                </div>
            </div>
            <div className="cards">

                {cities.map((city) => (
                    <div className="card">
                        <div className="card_img">
                            <img src={city.img} alt="" />
                        </div>
                        <div className="card_content">
                            <div className="card_title">{city.name}</div>
                            <div className="card_subtitle">{city.count}</div>
                        </div>
                    </div>
                ))}
               
            </div>
            
        </>
    )
}