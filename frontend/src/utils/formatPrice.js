const formatPrice = (priceString) => {
  if (!priceString) return '';

  // Remove ' TL' suffix and ensure dot as decimal separator
  let cleanedPrice = priceString.replace(' TL', '').replace(',', '.');

  // Parse as a float
  const price = parseFloat(cleanedPrice);

  if (isNaN(price)) {
    return priceString; // Return original if not a valid number
  }

  // Format to two decimal places and add ' TL' suffix
  return `${price.toFixed(2)} TL`;
};

export default formatPrice;
