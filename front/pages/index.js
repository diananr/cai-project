import React, { useState } from 'react';
import Head from 'next/head';

export default function Home() {
  const [form, setForm] = useState({
    age: 0,
    work: 0,
    link: 0,
    alcohol: 0,
    smoke: 0,
    drugs: 0,
    addiction: 0,
    risk: 0,
  });
  const [violence, setViolence] = useState(0);

  const handleParam = () => (e) => {
    const name = e.target.name;
    const value = e.target.value;
    setForm((prevState) => ({
      ...prevState,
      [name]: Number(value),
    }));
  };

  const formSubmit = async (e) => {
    e.preventDefault();
    const res = await fetch('http://localhost:8080/prediction', {
      method: 'POST',
      body: JSON.stringify(form),
      mode: 'cors',
      headers: new Headers(),
    });
    const data = await res.json();
    setViolence(data.Violence);
    document.getElementById('modal').classList.add('show', 'modalStyles');
  };

  const resetModal = () => {
    setViolence(0);
    document.getElementById('modal').classList.remove('show', 'modalStyles');
    setForm({
      age: 0,
      work: 0,
      link: 0,
      alcohol: 0,
      smoke: 0,
      drugs: 0,
      addiction: 0,
      risk: 0,
    });
  };

  return (
    <>
      <Head>
        <title>Predicción - CAI</title>
        <link rel='icon' href='/favicon.ico' />
        <link
          href='https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css'
          rel='stylesheet'
          integrity='sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC'
          crossOrigin='anonymous'
        />
        <link rel='stylesheet' href='/styles.css' />
      </Head>
      <main>
        <h1 className='title'>Predicción de CAI</h1>
        <form className='form' onSubmit={formSubmit}>
          <div className='item'>
            <label>Edad</label>
            <input type='number' name='age' required placeholder='Edad' className='form-control' value={form.age} onChange={handleParam()} />
          </div>
          <div className='item'>
            <label>Trabaja</label>
            <input type='number' name='work' required placeholder='Trabaja' className='form-control' value={form.work} onChange={handleParam()} />
          </div>
          <div className='item'>
            <label>Vínculo</label>
            <input type='number' name='link' required placeholder='Vinculo' className='form-control' value={form.link} onChange={handleParam()} />
          </div>
          <div className='item'>
            <label>Alcohol</label>
            <input
              type='number'
              name='alcohol'
              required
              placeholder='Alcohol'
              className='form-control'
              value={form.alcohol}
              onChange={handleParam()}
            />
          </div>
          <div className='item'>
            <label>Fuma</label>
            <input type='number' name='smoke' required placeholder='Fuma' className='form-control' value={form.smoke} onChange={handleParam()} />
          </div>
          <div className='item'>
            <label>Drogas</label>
            <input type='number' name='drugs' required placeholder='Drogas' className='form-control' value={form.drugs} onChange={handleParam()} />
          </div>
          <div className='item'>
            <label>Adicción</label>
            <input
              type='number'
              name='addiction'
              required
              placeholder='Adicción'
              className='form-control'
              value={form.addiction}
              onChange={handleParam()}
            />
          </div>
          <div className='item'>
            <label>Riesgo</label>
            <input type='number' name='risk' required placeholder='Riesgo' className='form-control' value={form.risk} onChange={handleParam()} />
          </div>
          <button type='submit' className='btn btn-primary button'>
            Enviar
          </button>
        </form>
      </main>

      <div id='modal' className='modal fade' tabIndex='-1' aria-labelledby='exampleModalLabel' aria-hidden='true'>
        <div className='modal-dialog'>
          <div className='modal-content'>
            <div className='modal-header'>
              <h5 className='modal-title' id='exampleModalLabel'>
                Resultado
              </h5>
              <button
                type='button'
                className='btn-close'
                data-bs-dismiss='modal'
                aria-label='Cerrar'
                onClick={() => {
                  resetModal();
                }}
              ></button>
            </div>
            <div className='modal-body'>El tipo de violencia es: {violence}</div>
            <div className='modal-footer'>
              <button
                type='button'
                className='btn btn-secondary'
                data-bs-dismiss='modal'
                onClick={() => {
                  resetModal();
                }}
              >
                Cerrar
              </button>
            </div>
          </div>
        </div>
      </div>
      <script
        src='https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js'
        integrity='sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM'
        crossOrigin='anonymous'
      />
    </>
  );
}
