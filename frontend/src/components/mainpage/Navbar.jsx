import { useState } from 'react';

export function Navbar() {
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const [isRegisterModalOpen, setIsRegisterModalOpen] = useState(false);

  
  const openLoginModal = () => {
    setIsLoginModalOpen(true);
  };

  const closeLoginModal = () => {
    setIsLoginModalOpen(false);
  };

  const openRegisterModal = () => {
    setIsLoginModalOpen(false);
    setIsRegisterModalOpen(true);
  };

  const closeRegisterModal = () => {
    setIsRegisterModalOpen(false);
  };

  const handleLoginButtonClick = () => {
    closeRegisterModal();
    openLoginModal();
  };

 const handleBackButtonClick = () => {
    closeRegisterModal();
    openLoginModal();
  };
 
    return(
        <nav className="navbar_mainpage">
            <div className="nav_links">
                <ul className="nav_links_block1">
                    <li><a href="/city">Города</a></li>
                    <li><a href="/hotel">Отели</a></li>
                </ul>
                <ul className="nav_links_block2">
                <li><a className="nav_link1" href="/create_request">Сдать жилье</a></li>
                <li><a className="nav_link2" href="#" onClick={openLoginModal}>Войти</a></li>
                </ul>
            </div>

            <div className="nav_titles">
                <div className="nav_title">Бронируй комфортное жилье с нами</div>
                <div className="nav_subtitle">Найди лучшее предложение для себя</div>
            </div>
            <div className="search_area">
                <div className="row">
                    <div className="column">
                        <div className="search_area_title">Город</div>
                        { <input type="text" placeholder="Выберите метоположение" /> }
                        {/* <label for="city">Город</label> */}
                    
                    </div>
                        <div className="column">
                            <div className="search_area_title">Дата заезда - Дата отъезда</div>
                            {/* <input type="text" placeholder="Выберите дату" /> */}
                            {/* <input type="date" id="date" name="date" min="2023-06-05"></input>
                            
                            <input type="date" id="date" name="date" min="2023-06-05"></input> */}
                            <input type="text" name="daterange" value="01/01/2023 - 01/12/2023" />
                        </div>
                        <div className="column">
                            <div className="search_area_title">Гости</div>

                        </div>
                        <button className="search_btn">Найти</button>
                        </div>
                    </div>
            {isLoginModalOpen && (
        <div className="modal-overlay">
          <div className="modal">
          <div className="enter">
          <button className="close-button"  onClick={closeLoginModal} ><img src="./src/assets/img/navbar/close.svg" alt="" /></button>
          </div>
            <div className="modal-header">
                <h2>Вход в личный кабинет</h2>
            </div>
            <div className="modal-content">
              <div className="form-group">
                <label htmlFor="email">E-mail</label>
                <input className="input" type="email" id="email" name="email"  placeholder="Введите адрес электронной почты"/>
              </div>
              <div className="form-group">
                <label htmlFor="password">Пароль</label>
                <input className="input" type="password" id="password" name="password" placeholder="Введите пароль" />
              </div>
              <button className="login-button">Войти</button>
              <div className="form-output">
              <a>У вас еще нет аккаунта?</a>
              <button className="register-button-next" onClick={openRegisterModal}>Зарегистрироваться</button>
              </div>
            </div>
          </div>
        </div>
      )}

        {isRegisterModalOpen && (
        <div className="modal-overlay">
            <div className="modal">
                <div className="button-back-close">
                <button className="back-button" onClick={handleBackButtonClick}><img src="./src/assets/img/navbar/back.svg" alt="" />Назад</button>
                <button className="close-button" onClick={closeRegisterModal}><img src="./src/assets/img/navbar/close.svg" alt="" /></button>
                </div>
                <div className="modal-header_registration">
                    <h2>Регистрация</h2>
                </div>
                    <div className="modal-content">
                        <div className="form-group">
                        <label htmlFor="name">Введите имя</label>
                        <input className="input" type="text" id="name" name="name" placeholder="Имя" />
                        </div>
                        <div className="form-group">
                        <label htmlFor="email">E-mail</label>
                        <input className="input" type="email" id="email" name="email" placeholder="Введите адрес электронной почты" />
                        </div>
                        <div className="form-group">
                        <label htmlFor="password">Пароль</label>
                        <input className="input" type="password" id="password" name="password" placeholder="Введите пароль"/>
                        </div>
                        <button className="register-button">Зарегистрироваться</button>
                    </div>
            </div>
        </div>
      )}
    </nav>
    )
}