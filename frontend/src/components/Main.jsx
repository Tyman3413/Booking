import {Navbar} from './mainpage/Navbar.jsx'
import {PopDir} from './mainpage/PopDir.jsx'
import { PopHotel } from './mainpage/PopHotel.jsx'
import {Create} from './mainpage/Create.jsx'
import { SubNews } from './mainpage/SubNews.jsx'
import { Footer } from './mainpage/Footer.jsx'

export function Main(){
    return(
        <>
            <header>
                <Navbar/>
            </header>
            <body>
                <PopDir/>
                <PopHotel/>
                <Create/>
                <SubNews/>
                <Footer/>
            </body>
            
        </>
        
    )
}