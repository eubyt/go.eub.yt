import { useState } from 'react';

function validURL(str) {
  const pattern = new RegExp('^(https?:\\/\\/)?' + // protocol
    '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|' + // domain name
    '((\\d{1,3}\\.){3}\\d{1,3}))' + // OR ip (v4) address
    '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*' + // port and path
    '(\\?[;&a-z\\d%_.~+=-]*)?' + // query string
    '(\\#[-a-z\\d_]*)?$', 'i'); // fragment locator
  return !!pattern.test(str);
}

function protocolExist(url) {
  return url.startsWith('http://') || url.startsWith('https://');
}

function StateInitial({ setState, setUrlShort, setAlert }) {
  const [url, setUrl] = useState('');

  async function onClickShort(e) {
    let newUrl = url;
    e.preventDefault();
    setAlert('');

    // TODO adicionar alerta de URL inválida
    if (!validURL(url)) {
      setAlert('URL inválida');
      return;
    }

    // Adicionar protocolo se não existir
    if (!protocolExist(url)) {
      newUrl = `https://${url}`
    }

    setState('loading');
    const api = await fetch('/api/shorten', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ newUrl }),
    })
    const data = await api.json();
    if (api.status === 201) {
      setUrlShort(`https://eub.yt/${data.code}`);
      setState('shortened');
      return;
    }
    setState('initial');
  }

  return <>
    <form className="mt-6">
      <div className="pl-4 pr-4 rounded-md shadow-sm flex items-center">
        <input
          id="url"
          type="text"
          placeholder="URL do site que deseja encurtar"
          required
          autoComplete="off"
          className="py-4 rounded-md w-full text-gray-500 focus:outline-none focus:shadow-outline border-transparent text-xl font-normal"
          onChange={(e) => setUrl(e.target.value)}
        />
      </div>
      <div className="mt-6">
        <span className="block w-full rounded-md shadow-sm">
          <button
            type="submit"
            onClick={onClickShort}
            className="w-full flex justify-center py-2 px-4 border border-transparent text-xx font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out"
          >
            Shorten
          </button>
        </span>
      </div>
    </form>
  </>
}

function StateLoading() {
  return <>
    <div className="text-center">
      <div className="text-xl font-bold">
        Shortening...
      </div>
    </div>
  </>
}

function StateShortened({ urlShort }) {
  return <>
    <div className="text-center">
      <div className="pl-4 pr-4 rounded-md shadow-sm flex items-center">
        <input
          type="text"
          className="py-4 rounded-md w-full text-gray-500 focus:outline-none focus:shadow-outline border-transparent text-xl font-normal text-center"
          value={urlShort}
          disabled
        />
      </div>
      <div className="mt-6 text-xl font-bold">
        <button
          className="w-full flex justify-center py-2 px-4 border border-transparent text-xx font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700 transition duration-150 ease-in-out"
          onClick={() => navigator.clipboard.writeText(urlShort)}
        >
          Copiar
        </button>
      </div>
    </div>
  </>
}

function App() {
  const [urlShort, setUrlShort] = useState('example.com');
  const [state, setState] = useState('initial');
  const [alert, setAlert] = useState('');

  return (
    <div className="application">
      <div className="relative flex min-h-screen flex-col justify-center overflow-hidden md:bg-gray-50 py-6 sm:py-12">
        <div className="relative bg-white px-6 pt-10 pb-8 md:shadow-xl ring-1 sm:mx-auto sm:max-w-lg sm:rounded-lg sm:px-10 ring-transparent">
          <div className="text-center">
            <h2 className="text-3xl leading-9 font-extrabold tracking-tight text-gray-900 sm:text-4xl sm:leading-10">
              go.eub.yt
            </h2>
            <p className="mt-3 text-lg leading-6 text-gray-500">
              Apenas um encurtador de links simples e rápido.
            </p>
            {alert && <div className="mt-3 text-sm leading-5 text-red-600">{alert}</div>}
          </div>

          <div className="mt-6">
            <div className="mt-6">
              {state === 'initial' && <StateInitial setState={setState} setUrlShort={setUrlShort} setAlert={setAlert} />}
              {state === 'loading' && <StateLoading />}
              {state === 'shortened' && <StateShortened urlShort={urlShort} />}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
