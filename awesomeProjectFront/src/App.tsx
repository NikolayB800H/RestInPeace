import { Routes, Route, Navigate } from 'react-router-dom';
import { DataTypes } from './pages/DataTypes'
import { DataTypeInfo } from './pages/DataTypeInfo'
import { ForecastApps } from './pages/ForecastApps'
import NavigationBar from './components/NavBar';

function App() {

  return (
    <>
      <NavigationBar />
      <div className='container-xl px-2 px-sm-3'>
        <Routes>
          <Route path="/" element={<Navigate to="data_types" />} />
          <Route path="/data_types" element={<DataTypes />} />
          <Route path="/data_types/:data_type_id" element={<DataTypeInfo />} />
          <Route path="/forecast_applications" element={<ForecastApps />} />
        </Routes>
      </div>
    </>
  )
}

export default App