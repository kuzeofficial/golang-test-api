async function getUsers() {
  return await fetch("http://localhost:3000/users");
}
const App = () => {
  return (
    <div>
      <button onClick={() => getUsers()}>Fetch Users</button>
    </div>
  );
};

export default App;
