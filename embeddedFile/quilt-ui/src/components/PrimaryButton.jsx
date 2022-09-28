import React from 'react';
import Button from '@material-ui/core/Button';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/core/styles';

import { typography, colors } from '../styles/globalStyles';

const useStyles = makeStyles(() => ({
    primaryButton: {
        width: "10rem",
        height: "3rem",
        borderRadius: "25px",
        backgroundColor: colors.primary01Color100,
        ...typography.actionText
    }
}));

const PrimaryButton = ({onClick, buttonText}) => {
    const classes = useStyles();
    
    return (<Button
      variant="contained" 
      color="primary"
      className={clsx(classes.primaryButton)} 
      onClick={onClick}
    >
      {buttonText}
  </Button>)
}

export default PrimaryButton;