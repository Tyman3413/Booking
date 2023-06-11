export function Footer(){
    return(
        <>
            <div className="footer">
                <div className="container">
                    <div className="footer_header">
                        <div className="row_footer">
                            <div className="column_footer">
                                <div className="footer_number">+7 (978) 000 0000</div>
                                <div className="footer_email">support@booking.com</div>
                            </div>
                            <div className="column_footer">
                                <div className="links">
                                    <div className="link"><a href="#">О проекте</a></div>
                                    <div className="link"><a href="#">Отзывы</a></div>
                                    <div className="link"><a href="#">Вопросы и ответы</a></div>
                                    <div className="link"><a href="#">Контакты</a></div>
                                    <div className="link"><a href="#">Поддержка</a></div>
                                </div>
                            </div>
                            <div className="column_footer">
                                <div className="links">
                                    <div className="link"><a href="#">Разместить объект</a></div>
                                    <div className="link"><a href="#">Помощь</a></div>
                                </div>
                            </div>
                            <div className="column_footer">
                                <div className="brands">
                                    <div className="brand"><img src="src/assets/img/brands/visa.png" alt="" /></div>
                                    <div className="brand"><img src="src/assets/img/brands/mastercard.png" alt="" /></div>
                                    <div className="brand"><img src="src/assets/img/brands/mir.png" alt="" /></div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="line"></div>
                    <div className="footer_bottom">
                        <div className="row_footer">
                            <div className="footer_bottom_element_logo">© 2023</div>
                            <div className="footer_bottom_element"><a href="#">Пользовательское соглашение</a></div>
                            <div className="footer_bottom_element"><a href="#">Политика конфиденциальности</a></div>
                            <div className="footer_bottom_element"><a href="#">Договор оферты</a></div>
                        </div>
                    </div>
                </div>
                <div id="up" className="up_btn"><img src="src/assets/img/components/up_arrow.png" alt="" /></div>
            </div>
        </>
    )
}
