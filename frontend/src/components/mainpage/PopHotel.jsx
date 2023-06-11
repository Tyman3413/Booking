import React, { useRef, useState } from 'react';

const hotels = [
  {
    img: './src/assets/img/hotels/1.png',
    name: 'Парк-отель «Песочная бухта»',
    rating: 4.1,
    address: 'Севастополь, ул. Ефремова, д. 38',
    price: 'От 5 700 ₽',
  },
  {
    img: './src/assets/img/hotels/2.png',
    name: 'Отель «Бристоль»',
    rating: 4.3,
    address: 'Ялта, ул. Рузвельта, д. 10',
    price: 'От 4 250 ₽',
  },
  {
    img: './src/assets/img/hotels/3.png',
    name: 'Отель «Петр»',
    rating: 4.4,
    address: 'Евпатория, ул. Гоголя, д. 15а',
    price: 'От 4 860 ₽',
  },
  {
    img: './src/assets/img/hotels/3.png',
    name: 'Отель «Петр»',
    rating: 4.4,
    address: 'Евпатория, ул. Гоголя, д. 15а',
    price: 'От 4 860 ₽',
  },
  {
    img: './src/assets/img/hotels/3.png',
    name: 'Отель «Петр»',
    rating: 4.4,
    address: 'Евпатория, ул. Гоголя, д. 15а',
    price: 'От 4 860 ₽',
  },
];

export function PopHotel() {
  const sliderRef = useRef(null);
  const [currentIndex, setCurrentIndex] = useState(0);

  const nextSlide = () => {
    const slideCount = hotels.length;
    setCurrentIndex((prevIndex) => (prevIndex + 1) % slideCount);
  };

  const prevSlide = () => {
    const slideCount = hotels.length;
    setCurrentIndex((prevIndex) => (prevIndex - 1 + slideCount) % slideCount);
  };

  const visibleHotels = hotels.slice(currentIndex, currentIndex + 3);

  return (
    <>
      <div className="section_title">
        <div className="name_title">Популярные отели</div>
        <div className="more">
          <div className="more_title">Смотреть больше</div>
          <div className="more_img">
            <img src="./src/assets/img/components/arrow_more.png" alt="" />
          </div>
        </div>
      </div>

      <div className="slider" ref={sliderRef}>
        <div className="cards">
          {visibleHotels.map((hotel, index) => (
            <div key={index} className="card">
              <div className="card_img">
                <div className="card_hotel_img">
                  <img src={hotel.img} alt="" />
                </div>
                <div className="card_like">
                  <img src="./src/assets/img/components/like.png" alt="" />
                </div>
              </div>
              <div className="card_content">
                <div className="card_content_header">
                  <div className="hotel_title">{hotel.name}</div>
                  <div className="card_rating_block">
                    <div className="card_rating">{hotel.rating}</div>
                    <div className="card_rating_img">
                      <img src="./src/assets/img/components/star.png" alt="" />
                    </div>
                  </div>
                </div>
                <div className="card_address">{hotel.address}</div>
                <div className="card_price">{hotel.price}</div>
              </div>
            </div>
          ))}
        </div>

        <div className="slider_buttons">
          <button className="left-button" onClick={prevSlide}><img src="./src/assets/img/navbar/left.svg" alt="" /></button>
          <button className="right-button" onClick={nextSlide}><img src="./src/assets/img/navbar/right.svg" alt="" /></button>
        </div>
      </div>
    </>
  );
}
