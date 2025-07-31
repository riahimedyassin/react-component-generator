import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
{{ROUTE_IMPORT}}

export default function {{FILE_NAME}}() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        {{ROUTES}}
      </Routes>
    </Router>
  );
}
