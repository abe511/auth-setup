import "./button-style.css";


type ButtonProps = {
  id?: string,
  className?: string,
  text?: string,
  type?: "submit" | "reset" | "button" | undefined
}

const Button = ({id, className, text, type}: ButtonProps) => {
  return (
    <button id={id} className={`btn ${className}`} type={type}>{text}</button>
  )
};

export default Button;